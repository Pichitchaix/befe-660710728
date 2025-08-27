package main

import (
	"github.com/gin-gonic/gin" /*Gin FrameWork ไม่ใช่ Library*/
)

type User struct { /*Struct รูปแบบ json*/
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	/*End Point คือ URL ที่โปรแกรมหรือแอปพลิเคชันใช้ในการติดต่อส่งข้อมูลระหว่างกัน*/
	r.GET("/users", func(c *gin.Context) {
		user := []User{{ID: "1", Name: "Pichitchai"}}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
