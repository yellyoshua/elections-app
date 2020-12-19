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
		{"PORT", "3000"},
		{"DATABASE_URI", "mongodb://root:dbpwd@localhost:27017"},
		{"DATABASE_NAME", "golang"},
	}

	for _, env := range envs {
		if value := os.Getenv(env.name); len(value) == 0 {
			os.Setenv(env.name, env.value)
		}
	}
	return
}
