package pkg

import (
	"os"
)

type AppConfig struct {
	AWSConfig *AWSConfig
	DBConfig  *DBConfig
}

type AWSConfig struct {
	AccessKeyID string
	SecretKey   string
	Region      string
}

type DBConfig struct {
	Location string
	Username string
	Password string
	DBName   string
}

func LoadAppConfig() *AppConfig {
	awsConfig := loadAWSConfig()
	DBConfig := loadDBConfig()
	return &AppConfig{
		AWSConfig: awsConfig,
		DBConfig:  DBConfig,
	}

}

func loadDBConfig() *DBConfig {
	location := os.Getenv("DB_URL")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return &DBConfig{
		Location: location,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
}

func loadAWSConfig() *AWSConfig {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")
	return &AWSConfig{
		AccessKeyID: accessKeyID,
		SecretKey:   secretKey,
		Region:      region,
	}
}
