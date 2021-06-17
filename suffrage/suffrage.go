package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"github.com/yellyoshua/elections-app/commons/constants"
	"github.com/yellyoshua/elections-app/commons/logger"
	"github.com/yellyoshua/elections-app/commons/models/ballot_models"
)

func main() {
	setupEnvironments()

	port := os.Getenv("PORT")
	if noPort := len(port) == 0; noPort {
		port = "3000"
	}

	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("RABBITMQ_URI")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()
	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		constants.QueueSuffrageServiceName, // queue name
		true,                               // durable
		false,                              // auto delete
		false,                              // exclusive
		false,                              // no wait
		nil,                                // arguments
	)
	if err != nil {
		panic(err)
	}

	exit := make(chan bool)
	go handleInterrupt(exit)

	// Create a new Fiber instance.
	app := gin.New()

	// HTTP SERVER
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        app,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Add route for send message to Service 1.
	app.GET("/send", func(c *gin.Context) {
		var ballot = ballot_models.BallotModel{}
		json.Marshal(ballot)

		// Create a message to publish.
		message := amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(c.Query("msg")),
		}

		// Attempt to publish a message to the queue.
		if err := channelRabbitMQ.Publish(
			"",                                 // exchange
			constants.QueueSuffrageServiceName, // queue name
			false,                              // mandatory
			false,                              // immediate
			message,                            // message to publish
		); err != nil {
			c.AbortWithStatus(500)
		} else {
			c.Status(200)
			c.Done()
		}

	})
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Panic("Error trying start gin-gonic server -> " + err.Error())
		}
	}()

	<-exit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		logger.Panic("Gracefull HTTP server shutdown failed: " + err.Error())
	}
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

func handleInterrupt(exit chan bool) {
	ch := make(chan os.Signal)
	// the channel used with signal.Notify should be buffered (SA1017)
	signal.Notify(ch, os.Interrupt)
	<-ch

	logger.Info("Closing server")
	exit <- true
}
