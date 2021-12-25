package main

import (
	"pkg/coba/config"
	"pkg/coba/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func status(con *gin.Context) {
// 	con.JSON(200, gin.H{
// 		"message": "Success",
// 	})
// }

var (
	db             *gorm.DB                  = config.DatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	route := gin.Default()
	route.SetTrustedProxies(nil)
	// route.GET("/", status)

	// Make group of route api/auth
	authRoute := route.Group("api/auth") 
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}
	route.Run()
}
