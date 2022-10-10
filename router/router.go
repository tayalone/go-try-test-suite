package router

import (
	"net/http"

	"go-try-test-suite/geo"

	"github.com/gin-gonic/gin"
)

func getDesc(g geo.Geo) geo.Desc {
	return geo.Desc{
		Area:      g.Area(),
		Perimeter: g.Perimeter(),
	}
}

/*SetUpRouter return gin.Engine Prepare for intregation test*/
func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	type circleBody struct {
		R float64 `json:"r" binding:"required"`
	}

	r.POST("/circle", func(c *gin.Context) {
		var body circleBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		desc := getDesc(&geo.Circle{R: body.R})

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"desc":    desc,
		})
	})
	return r
}
