package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ua/com/booking/domain"
)

type database []domain.Room

func (db database) getRoomById(id int) domain.Room {
	for _, v := range db {
		if v.Id == id {
			return v
		}
	}

	return domain.Room{}
}

func main() {
	db := database{
		{
			1,
			"room 0345",
			"some descripiton",
			[]domain.Slot{},
		},
		{
			2,
			"room 0455",
			"some other descripiton",
			[]domain.Slot{},
		},
	}

	router := gin.Default()

	router.GET("room/:id", func(c *gin.Context) {
		id := c.Param("id")
		num, _ := strconv.Atoi(id)
		c.String(http.StatusOK, "This is a room: %#v", db.getRoomById(num))
	})
	router.Run()
}
