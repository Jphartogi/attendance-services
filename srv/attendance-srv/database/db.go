package database

import (
	"log"
	"strings"

	"github.com/Jphartogi/attendance-services/srv/attendance-srv/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/spf13/viper"
)

var db *pg.DB

// InitDB ...
func InitDB() *pg.DB {
	stringSlice := []string{viper.GetString("dbConfig.address"), viper.GetString("dbConfig.port")}
	DBAddress := strings.Join(stringSlice, ":")
	opts := &pg.Options{
		User:     viper.GetString("dbConfig.username"),
		Password: viper.GetString("dbConfig.password"),
		Addr:     DBAddress,
		Database: viper.GetString("dbConfig.dbName"),
	}
	db = pg.Connect(opts)
	if db == nil {
		log.Fatal("Failed to connect to DB")
	}
	log.Printf("Connected to db")
	createEmployeeTable(db)
	return db
}

// GetDB for getting db instances
func GetDB() *pg.DB {
	return db
}

// createProjectsTable create Table for user who is a maker
func createEmployeeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.Employee{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating users table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Employee table created")
	return nil
}
