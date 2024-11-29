package routes

import (
	"api/src/controllers" // ajuste o caminho conforme necessário

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// ... outras rotas ...

	router.POST("/login", controllers.LoginSuperAdmin(db)) // Rota para login do superadmin
}