package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type AppConfig struct {
	ServerPort			string
	Dsn					string
	AppSecret			string
	TwilioAccountSid	string
	TwilioAuthToken		string
	TwilioPhoneNumber	string
}

func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("DSN is not set")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("APP_SECRET is not set")
	}


	return AppConfig{ServerPort: httpPort, Dsn: dsn, AppSecret: appSecret,
		TwilioAccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioPhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),	
	}, nil
}