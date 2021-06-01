package main

import (
	"log"

	"github.com/gabrielpbzr/findme/infra"
	"github.com/gin-gonic/gin"
)

func main() {
	dbHandler := infra.Database{DSN: "data/data.db"}

	db, err := dbHandler.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer dbHandler.Close(db)
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "HELLO GIN-GONIC!")
	})

	router.Run()
}
