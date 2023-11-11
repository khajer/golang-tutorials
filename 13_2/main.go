package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DB_NAME     = "members"
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_HOST     = "localhost"
)

type Member struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func init() {
	fmt.Println("Init() like a constructor")
	var err error

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Connected to the database")
}

func GetMembers(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, email FROM members")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch members"})
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan members"})
		}
		members = append(members, member)
	}

	c.JSON(http.StatusOK, members)

}

func main() {
	r := gin.Default()
	r.GET("/members", GetMembers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
