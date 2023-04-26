package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
 App         string // project name
 Environment string // development, test, prod
 Version     string //

 ServiceHost string // localhost
 HTTPPort    string //7777

 PostgresHost     string // localhost
 PostgresPort     int    //5432
 PostgresUser     string // admin
 PostgresPassword string // just_db
 PostgresDatabase string // just_db

 DefaultOffset string // 0
 DefaultLimit  string // 10
}

// Load ...
func Load() Config {
 if err := godotenv.Load(".env"); err != nil {
  fmt.Println("No .env file found")
 }

 config := Config{}

 // type assertation
 config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "todo"))
 config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev"))
 config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

 config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
 config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

 config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
 config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
 config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "islombek"))
 config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "lib1234"))
 config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABgit initASE", "todo"))

 config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
 config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

 return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
 val, exists := os.LookupEnv(key)

 if exists {
  return val
 }

 return defaultValue
}