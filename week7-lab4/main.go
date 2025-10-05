package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// อ่าน environment variable
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// global db
var db *sql.DB

func initDB() {
	var err error

	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	if host == "" || port == "" || user == "" || password == "" || name == "" {
		log.Fatal("Database environment variables not set properly")
	}

	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_NAME:", name)
	fmt.Println("DB_USER:", user)
	fmt.Println("DB_PASSWORD:", password)
	fmt.Println("DB_PORT:", port)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)
	
	fmt.Println(connStr)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to database")
}

func main() {
	initDB()
	defer db.Close() // ปิด database connection ตอนโปรแกรม exit

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	r.Run(":8080")
}
