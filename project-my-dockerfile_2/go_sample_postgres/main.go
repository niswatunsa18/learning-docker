package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func main() {
// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.HTML(http.StatusOK, "Hello, Docker! <3")
// 	})

// 	httpPort := os.Getenv("PORT")
// 	if httpPort == "" {
// 		httpPort = "80"
// 	}

// 	e.Logger.Fatal(e.Start(":" + httpPort))
// }

func main() {
	// Get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to database
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	// postgresql://username:password@database:5432/postgresdb
	fmt.Println("Connect to: ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed: ", err)
	}

	// test ping to database
	if err := db.Ping(); err != nil {
		fmt.Printf("Connetion to database failed (DB_HOST: %s): %s\n", dbHost, err)
	} else {
		fmt.Println("Successfully connected to database: ", db)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Ouch!")
	})
	e.GET("/ping", func(c echo.Context) error {
		_, err = db.Exec("INSERT INTO timestamp (timestamp) VALUES ($1)", time.Now())
		if err := db.Ping(); err != nil {
			fmt.Println("Ping Failed")
			return c.HTML(http.StatusBadRequest, "Ouch!")
		}
		fmt.Println("Ping Successfull")
		return c.HTML(http.StatusOK, "PONG! \n")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "78"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
