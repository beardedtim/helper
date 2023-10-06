package datastore

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Datastore struct {
	Database *gorm.DB
}

var DatastoreInstance = Datastore{}

func (datastore *Datastore) Connect() {
	log.Trace("Connecting to Database")

	if datastore.Database != nil {
		log.Warn("Trying to connect when we already have a connection. This will be a noop but may become a panic in future releases")
		return
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		log.WithError(err).Panic("We could not connect to the database")
	}

	datastore.Database = db
	log.Trace("Database Connected")
}

func (datastore *Datastore) Disconnect() {
	log.Trace("Disconnecting from Database")
	if datastore.Database == nil {
		log.Warn("Trying to disconnect without connecting. This will be a noop for now but could cause a panic in future releases")
		return
	}

	db, err := datastore.Database.DB()

	if err != nil {
		log.Warn("There was an issue trying to grab the raw db connection")
		return
	}

	db.Close()

	log.Trace("Database Disconnected")
}

func (datastore *Datastore) IsHealthy() bool {
	log.Trace("Checking if we have a healthy connection to the Database")
	db, err := datastore.Database.DB()

	if err != nil {
		log.WithError(err).Warn("There was an error checking if the datastore is healthy")
		return false
	}

	err = db.Ping()

	if err != nil {
		log.WithError(err).Warn("There was an error checking if the datastore is healthy")
		return false
	}

	log.Trace("We have a healthy connection to the Database")

	return true
}

// If you create a new Model, you need to add it here
func (datastore *Datastore) EnsureMigration() {
	datastore.Database.AutoMigrate(&UsersModel{})
	datastore.Database.AutoMigrate(&GroupsModel{})

}
