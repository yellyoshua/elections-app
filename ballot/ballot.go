package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"github.com/yellyoshua/elections-app/commons/constants"
	"github.com/yellyoshua/elections-app/commons/logger"
	"github.com/yellyoshua/elections-app/commons/models/ballot_models"
	"github.com/yellyoshua/elections-app/commons/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ballotRepoIndexes = []mongo.IndexModel{
	{
		Options: options.Index().SetName("institute _id").SetUnique(false).SetDefaultLanguage("en"),
		Keys:    bson.M{"institute": 1},
	},
	{
		Options: options.Index().SetName("endorsement _id").SetUnique(false).SetDefaultLanguage("en"),
		Keys:    bson.M{"endorsement": 1},
	},
	{
		Options: options.Index().SetName("suffrage _id").SetUnique(false).SetDefaultLanguage("en"),
		Keys:    bson.M{"suffrage": 1},
	},
}

func ballotRepository() repository.Client {
	databaseURI := os.Getenv("BALLOT_MONGODB_URI")
	databaseName := os.Getenv("BALLOT_MONGODB_DATABASE")

	return repository.New(repository.RepositoryConf{
		DatabaseURI:  databaseURI,
		DatabaseName: databaseName,
	}, bootstrapRepository)
}

func bootstrapRepository(db *mongo.Database) error {
	return nil
}

func main() {
	setupEnvironments()

	ballotRepository()

	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("RABBITMQ_URI")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		logger.Panic("error with rabbitmq channel -> %s", err)
		panic(err)
	}
	defer channelRabbitMQ.Close()

	queueSuffrage, errSuffrageQueue := channelRabbitMQ.QueueDeclare(
		constants.QueueSuffrageServiceName, // name
		false,                              // durable
		false,                              // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)

	if errSuffrageQueue != nil {
		logger.Panic("error declaring suffrage-queue -> %s", errSuffrageQueue)
	}

	// Subscribing to suffrage_service_queue for getting suffrages.
	suffrages, err := channelRabbitMQ.Consume(
		queueSuffrage.Name, // queue name
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no local
		false,              // no wait
		nil,                // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for suffrages")

	// Make a channel to receive suffrages into infinite loop.
	forever := make(chan bool)

	go func() {
		for suffrage := range suffrages {

			var ballot ballot_models.BallotModel
			if err := json.Unmarshal(suffrage.Body, &ballot); err == nil {
				log.Printf(" > Received message with correct struct!")
			}
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", string(suffrage.Body))
		}
	}()

	<-forever
}

// Environments if not exist .env file load system environments or defaults!
func setupEnvironments() {
	godotenv.Load(".env")

	envs := map[string]bool{
		"BALLOT_MONGODB_URI":      true,
		"BALLOT_MONGODB_DATABASE": true,
		"RABBITMQ_URI":            true,
	}

	for name, isRequired := range envs {
		value := os.Getenv(name)

		if existEnv := len(value) == 0; existEnv && isRequired {
			logger.Panic("%v environment variable is required.\n", name)
		}
	}
}
