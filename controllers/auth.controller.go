package controllers

import (
	"diplom_back/models"
	"diplom_back/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type AuthController struct {
	DB           *gorm.DB
	tokenService services.Jwt
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB: DB, tokenService: services.NewJWT()}
}

func (c *AuthController) SignUpUser(context *gin.Context) {
	var signUpUser models.SignUpUser

	if err := context.ShouldBindJSON(&signUpUser); err != nil {
		log.Fatal(err)
	}

	err := c.DB.Create(&models.User{
		Model:    gorm.Model{},
		Email:    signUpUser.Email,
		Username: signUpUser.Username,
		Password: signUpUser.Password,
	}).Error

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(200, gin.H{
		"message": "success",
	})
}

func (c *AuthController) SignInUser(ctx *gin.Context) {
	token := c.tokenService.GetNewAccessToken()
	ctx.JSON(200, token)
	/*var payload *models.SignInUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := c.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if result := utils.VerifyPassword(user.Password, payload.Password); result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")
	*/
	/*accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
	*/
	//ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
