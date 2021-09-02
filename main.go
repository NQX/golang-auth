package main

import ("github.com/gin-gonic/gin"
		"golang-auth/routes"
)

func main() {

	router := gin.Default()

	usersGroup := router.Group("users") {
		usersGroup.POST("register", routes.UsersRegister)
	}

	router.Run(":8080")
}
