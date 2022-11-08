package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	hostEnv := os.Getenv("MARIADB_HOST")
	if hostEnv == "" {
		log.Fatal("MARIADB_HOST must be set")
	}

	dbNameEnv := os.Getenv("MARIADB_DBNAME")
	if dbNameEnv == "" {
		log.Fatal("MARIADB_DBNAME must be set")
	}

	dbUserEnv := os.Getenv("MARIADB_USER")
	if dbUserEnv == "" {
		log.Fatal("MARIADB_USER must be set")
	}

	dbPassEnv := os.Getenv("MARIADB_PASS")
	if dbPassEnv == "" {
		log.Fatal("MARIADB_PASS must be set")
	}

	dbPortEnv := os.Getenv("MARIADB_PORT")
	dbPort, _ := strconv.Atoi(dbPortEnv)
	if dbPortEnv == "" {
		log.Fatal("MARIADB_PORT must be set")
	}

	dbConfig := DBConfig{
		Host:     hostEnv,
		Port:     dbPort,
		User:     dbUserEnv,
		Password: dbPassEnv,
		DBName:   dbNameEnv,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
