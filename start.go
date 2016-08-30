package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"strconv"
)

var dbmap = initDb()

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/rooms", GetRooms)
		v1.GET("/rooms/:id", GetRoom)
		v1.POST("/rooms", PostRoom)
		v1.PUT("/rooms/:id", UpdateRoom)
		v1.DELETE("/rooms/:id", DeleteRoom)
	}

	r.Run(":8080")
}

func GetRooms(c *gin.Context) {
	var rooms []Room
	_, err := dbmap.Select(&rooms, "SELECT * FROM room")

	if err == nil {
		c.JSON(200, rooms)
	} else {
		c.JSON(404, gin.H{"error": "no room(s) into the table"})
	}
}

func GetRoom(c *gin.Context) {

	id := c.Params.ByName("id")
	var room Room
	err := dbmap.SelectOne(&room, "SELECT * FROM room WHERE id=?", id)

	if err == nil {
		c.JSON(200, room)
	} else {
		c.JSON(404, gin.H{"error": "room not found"})
	}
}

func PostRoom(c *gin.Context) {
	var room Room
	c.Bind(&room)

	if room.Name != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO room (Name, Description) VALUES (?, ?)`, room.Name, room.Description); insert != nil {
			room_id, err := insert.LastInsertId()
			if err == nil {
				room.Id = room_id
				c.JSON(201, room)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	}

}

func UpdateRoom(c *gin.Context) {
	id := c.Params.ByName("id")
	var room Room
	err := dbmap.SelectOne(&room, "SELECT * FROM room WHERE id=?", id)

	if err == nil {
		var json Room
		c.Bind(&json)

		room_id, _ := strconv.ParseInt(id, 0, 64)

		room := Room{
			Id:          room_id,
			Name:        json.Name,
			Description: json.Description,
			Bookings:    json.Bookings,
		}

		if room.Name != "" {
			_, err = dbmap.Update(&room)
			if err == nil {
				c.JSON(200, room)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "room not found"})
	}
}

func DeleteRoom(c *gin.Context) {
	id := c.Params.ByName("id")

	var room Room
	err := dbmap.SelectOne(&room, "SELECT id FROM room WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&room)
		if err == nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			checkErr(err, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "room not found"})
	}

}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:golang@/booking")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Room{}, "Room").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")
	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
