package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"api/src/models"
	rest_error "api/src/configs" // ajuste o caminho conforme necessário
)

var jwtSecret = []byte("your_secret_key") // Substitua pela sua chave secreta

// Função para gerar o token JWT
func generateToken(user models.Users) (string, error) {
	claims := jwt.MapClaims{
		"uid":  user.Email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

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

		// Verifica se a senha está correta usando bcrypt
		if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(credentials.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, rest_error.NewUnauthorizedError("Senha não condiz com o email informado"))
			return
		}

		// Verifica se o tipo do usuário é SuperAdmin
		if user.Type == nil || *user.Type != "SuperAdmin" {
			c.JSON(http.StatusUnauthorized, rest_error.NewUnauthorizedError("Você não tem permissão para acessar esta área."))
			return
		}

		// Gera o token JWT
		token, err := generateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, rest_error.NewInternalServerError("Erro ao gerar token"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access-token": token,
			"token-type":  "Bearer",
			"client":      "ceJ7Fm6HfFsLrL4WonF6rw", // Substitua conforme necessário
			"expiry":      time.Now().Add(time.Hour * 24).Unix(), // Expiração do token
			"uid":         user.Email,
		})
	}
}