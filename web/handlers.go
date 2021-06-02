package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielpbzr/findme/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeHandlers(router *gin.Engine, positionService *core.PositionServiceDB) {
	// cria os handlers http
	router.GET("/", index)
	router.GET("/api/tracking/:id", findPosition(positionService))
	router.POST("/api/tracking", registerPosition(positionService))
}

func index(ctx *gin.Context) {
	ctx.String(200, "<h1>HELLO GIN-GONIC!</h1>")
}

func registerPosition(positionService *core.PositionServiceDB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		longitude, err := parseLongitude(ctx.PostForm("longitude"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		latitude, err := parseLatitude(ctx.PostForm("latitude"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		position := core.CreatePosition(longitude, latitude)
		err = positionService.Create(position)
		if err == nil {
			ctx.Status(http.StatusCreated)
			ctx.Header("Location", "/api/tracking/"+position.Id.String())
		}
	}
}

func findPosition(positionService *core.PositionServiceDB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := uuid.MustParse(ctx.Param("id"))
		position, err := positionService.Get(uuid)
		if err != nil {
			log.Println(err)
		}

		if position != nil {
			ctx.JSON(http.StatusOK, position)
			return
		}

		ctx.Status(http.StatusNotFound)
	}
}

func parseLongitude(strLongitude string) (float64, error) {
	const errorMsg = "longitude should be a value between -180 and 180"
	lon, err := strconv.ParseFloat(strLongitude, 64)
	if err != nil {
		return 0, fmt.Errorf(errorMsg)
	}

	if lon > 180 || lon < -180 {
		return 0, fmt.Errorf(errorMsg)
	}

	return lon, nil
}

func parseLatitude(strLatitude string) (float64, error) {
	const errorMsg = "latitude should be a value between -90 and 90"
	lat, err := strconv.ParseFloat(strLatitude, 64)
	if err != nil {
		return 0, fmt.Errorf(errorMsg)
	}

	if lat > 90 || lat < -90 {
		return 0, fmt.Errorf(errorMsg)
	}

	return lat, nil
}
