package controllers

import (
	rest_error "api/src/configs" // ajuste o caminho conforme necessário
	"api/src/models"             // ajuste o caminho conforme necessário
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Função para login de Superadmin
func LoginSuperAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Verifica se o corpo da requisição é válido
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, rest_error.NewBadRequestError("Dados inválidos"))
			return
		}

		// Verifica se o email e a senha estão preenchidos
		if credentials.Email == "" {
			c.JSON(http.StatusBadRequest, rest_error.NewBadRequestError("Email não pode ser vazio"))
			return
		}
		if credentials.Password == "" {
			c.JSON(http.StatusBadRequest, rest_error.NewBadRequestError("Senha não pode ser vazia"))
			return
		}

		var user models.Users
		// Verifica se o usuário existe
		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusUnauthorized, rest_error.NewUnauthorizedError("Email não encontrado"))
			} else {
				c.JSON(http.StatusInternalServerError, rest_error.NewInternalServerError("Erro ao acessar o banco de dados"))
			}
			return
		}

		// Verifica se a senha está correta
		if user.EncryptedPassword != credentials.Password { // Aqui você deve usar a função de comparação de hash se a senha estiver criptografada
			c.JSON(http.StatusUnauthorized, rest_error.NewUnauthorizedError("Senha não condiz com o email informado"))
			return
		}

		// Verifica se o tipo do usuário é SuperAdmin
		if user.Type == nil || *user.Type != "SuperAdmin" {
			c.JSON(http.StatusUnauthorized, rest_error.NewUnauthorizedError("Você não tem permissão para acessar esta área."))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido!", "user": user})
	}
}