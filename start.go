package main

import (
	"booking/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	db := domain.GetDatabase()
	router := gin.Default()

	router.GET("room/:id", func(c *gin.Context) {
		id := c.Param("id")
		num, _ := strconv.Atoi(id)
		c.String(http.StatusOK, "This is a room: %#v", db.GetRoomById(num))
	})
	router.Run()
}
