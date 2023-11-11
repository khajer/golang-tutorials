package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
	fmt.Println("Init like a constructor")
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

// Handler functions
func GetMembers(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, email FROM members")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch members"})
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan members"})
		}
		members = append(members, member)
	}

	return c.JSON(http.StatusOK, members)
}

// Start the Echo server
func main() {
	e := echo.New()

	// Routes
	e.GET("/members", GetMembers)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
