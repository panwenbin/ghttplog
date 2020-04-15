package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongodbUri        string
	MongodbDatabase   string
	MongodbCollection string
)

var Debug bool

func init() {
	checkEnv()
	LoadSetting()
}

func checkEnv() {
	_ = godotenv.Load()
	needChecks := []string{
		"MONGODB_URI",
		"MONGODB_DATABASE",
	}

	for _, envKey := range needChecks {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}

func LoadSetting() {
	MongodbUri = os.Getenv("MONGODB_URI")
	MongodbDatabase = os.Getenv("MONGODB_DATABASE")
	MongodbCollection = os.Getenv("MONGODB_COLLECTION")
	if MongodbCollection == "" {
		MongodbCollection = "http"
	}

	debug := os.Getenv("DEBUG")
	if debug != "" && debug != "false" && debug != "0" {
		Debug = true
	}
}
