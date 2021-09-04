package main

import (
	"context"
	"fmt"
	"golang-auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(dbMiddleware(*conn))

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("login", routes.UsersLogin)
	}

	router.Run(":8080")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5432/person")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}

	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
