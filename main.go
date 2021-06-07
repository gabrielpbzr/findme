package main

import (
	"log"

	"github.com/gabrielpbzr/findme/core"
	"github.com/gabrielpbzr/findme/infra"
	"github.com/gabrielpbzr/findme/web"
	"github.com/gin-gonic/gin"
)

func main() {
	dbHandler := infra.Database{DSN: "data/data.db"}

	db, err := dbHandler.Open()
	positionService := core.NewPositionService(db)

	if err != nil {
		log.Fatal(err)
	}

	defer dbHandler.Close(db)
	router := gin.Default()
	// Load template files
	router.LoadHTMLGlob("templates/*")
	// Initialize web handlers
	web.MakeHandlers(router, positionService)

	router.Run()
}
