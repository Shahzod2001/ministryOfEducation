package main

import (
	"ministryOfEducation/db"
	"ministryOfEducation/logger"
	"ministryOfEducation/pkg/controllers"
)

func main() {
	err := logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = controllers.RunRoutes()
	if err != nil {
		return
	}
}
