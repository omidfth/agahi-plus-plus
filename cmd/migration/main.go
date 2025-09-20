package main

import (
	"agahi-plus-plus/handler/app"
	"agahi-plus-plus/internal/constant"
	"agahi-plus-plus/internal/database"
	"agahi-plus-plus/internal/model"
	"flag"
	"gorm.io/gorm"
)

func main() {
	path := flag.String("e", constant.DefaultEnvPath, "env file path")
	flag.Parse()
	config, err := app.SetupViper(*path)
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(config.Database)
	if err != nil {
		panic(err)
	}
	migrate(db)
}

func migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(model.Post{})
	db.Debug().AutoMigrate(model.Config{})
	db.Debug().AutoMigrate(model.UserPayment{})
	db.Debug().AutoMigrate(model.Plan{})
	db.Debug().AutoMigrate(model.User{})
}
