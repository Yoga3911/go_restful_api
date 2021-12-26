package main

import (
	"pkg/coba/config"
	"pkg/coba/controller"
	"pkg/coba/repository"
	"pkg/coba/service"

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
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
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
