package repository

import (
	"context"
	"log"
	"os"
	"time"

	googleCloudStorage "cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type repo struct{}

var (
	db            *gorm.DB
	storageClient *googleCloudStorage.Client
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	loggerGorm := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Minute,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  false,        // Disable color
		},
	)

	dbOpen, errDBOpen := gorm.Open(postgres.Open(os.Getenv("DSN_SICERDAS")), &gorm.Config{
		Logger: loggerGorm,
	})

	if errDBOpen != nil {
		log.Panicln("error Connection Database")
	}
	db = dbOpen

	ctx := context.Background()
	sa := option.WithCredentialsFile("firebase.json")

	storageClient, err = googleCloudStorage.NewClient(ctx, sa)
	if err != nil {
		log.Fatalln(err)
	}
}
