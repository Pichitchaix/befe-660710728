package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"log"
	"github.com/gin-gonic/gin"
	 "net/http"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB(){
var err error
	
	host := getEnv("DB_HOST","")
	name := getEnv("DB_NAME","")
	user := getEnv("DB_USER","")
	password := getEnv("DB_PASSWORD","")
	port := getEnv("DB_PORT","")
	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_PORT:", port)
	fmt.Println("DB_USER:", user)
	fmt.Println("DB_PASSWORD:", password)
	fmt.Println("DB_NAME:", name)
	conSt :=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", host, port, user ,password ,name )
	// fmt.Println(conSt)
	db, err = sql.Open("postgres", conSt)
	if err != nil {
    log.Fatalf("failed to connect to database: %v", err)
}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to database",err)
	}

	log.Println("successfully connected to database")

}

func main(){
	initDB()
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy","error":err})
		// ถ้าใช้ http.StatusOK / http.StatusInternalServerError → import "net/http"
		// เช่น c.JSON(StatusInternalServerError, gin.H{"message": "unhealthy","error":err})
		return
	}

		c.JSON(200, gin.H{"message": "healthy"})
		
})
r.Run(":8080")
	}
