package setups

import (
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	name  string
	value string
}

// Environments if not exist .env file load system environments or defaults!
func Environments() {
	godotenv.Load(".env")

	envs := []environment{
		environment{"PORT", "3000"},
		environment{"DB_USER", "root"},
		environment{"DB_PASSWORD", "dbpwd"},
		environment{"DB_ADDR", "localhost"},
		environment{"DB_PORT", "27017"},
		environment{"DB_NAME", "golang"},
	}

	for _, env := range envs {
		if value := os.Getenv(env.name); len(value) == 0 {
			os.Setenv(env.name, env.value)
		}
	}
	return
}
