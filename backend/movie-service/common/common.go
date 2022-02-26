package common

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFileName   string `json:"logFileName"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	MongoAddrs      string `json:"mongoAddrs"`
	MongoDbName     string `json:"mongoDbName"`
	MongoDbUsername string `json:"mongoDbUsername"`
	MongoDbPassword string `json:"mongoDbPassword"`

	AuthAddr          string `json:"authAddr"`
	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
}

var Config *Configuration

const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000
	StatusMismatch    = 10
)

func LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting service logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFileName,
		MaxSize:    Config.LogMaxSize,
		MaxBackups: Config.LogMaxBackups,
		MaxAge:     Config.LogMaxAge,
	})
	log.SetLevel(log.DebugLevel)

	// Setting log formatter to json
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
