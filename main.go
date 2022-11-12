package main

import (
	"log"
	"os"

	"reach/reach-api/config"
	"reach/reach-api/models"
	"reach/reach-api/routes"

	"github.com/bamzi/jobrunner"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// MySQL
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(
		&models.CaseModel{},
		&models.DefinitionModel{},
		&models.InputSourceModel{},
		&models.OutputSourceModel{},
		&models.ReplyModel{},
		&models.SettingModel{},
		&models.TransformerModel{},
		&models.UserModel{})

	// Schedulers
	jobrunner.Start()
	//go jobrunner.Now(schedule.RabbitMQSchedule{})

	// jobrunner.Now(schedule.CollectDataSchedule{})
	// jobrunner.Schedule("@every 1d", schedule.CollectDividends{})
	// jobrunner.Now(schedule.CollectDividends{})
	//jobrunner.Now(schedule.Binance{})

	// Setup Router
	r := routes.SetupRouter()
	r.Run(":" + port)
}
