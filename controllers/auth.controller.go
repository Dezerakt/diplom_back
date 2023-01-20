package controllers

import (
	"diplom_back/models"
	"diplom_back/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB: DB}
}

func (c *AuthController) SignUpUser(context *gin.Context) {
	var signUpUser models.SignUpUser

	if err := context.ShouldBindJSON(&signUpUser); err != nil {
		utils.ErrorResponse(context, err)
	}

	err := c.DB.Create(&models.User{
		Model:    gorm.Model{},
		Email:    signUpUser.Email,
		Username: signUpUser.Username,
		Password: signUpUser.Password,
	}).Error

	if err != nil {
		utils.ErrorResponse(context, err)
	}

	token, err := bcrypt.GenerateFromPassword([]byte(signUpUser.Email), bcrypt.DefaultCost)

	if err != nil {
		utils.ErrorResponse(context, err)
	}

	var user models.User

	c.DB.Find(&user, "email = ?", signUpUser.Email)

	err = c.DB.Create(&models.Token{
		Model:  gorm.Model{},
		UserID: user.ID,
		Token:  string(token),
	}).Error

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(200, gin.H{
		"token": string(token),
	})
}

func (c *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	var user *models.User
	err := c.DB.First(&user, "email = ?", strings.ToLower(payload.Email)).Error
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	if result := utils.VerifyPassword(user.Password, payload.Password); result != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	token, err := bcrypt.GenerateFromPassword([]byte(payload.Email), bcrypt.DefaultCost)

	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	var firstToken *models.Token
	c.DB.First(&firstToken, "user_id = ?", user.ID)
	fmt.Println([]byte(token))

	err = c.DB.Model(&models.Token{}).Where("user_id = ?", &firstToken.UserID).Update("token", token).Error
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": string(token)})
}

func (c *AuthController) GetInfoByToken(ctx *gin.Context) {
	var (
		candidateToken models.Token
		candidateUser  models.User
	)

	type TokenArrived struct {
		Token string `json:"token"`
	}
	var arrivedToken *TokenArrived

	if err := ctx.ShouldBindJSON(&arrivedToken); err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	if err := c.DB.Where("token = ?", &arrivedToken.Token).First(&candidateToken).Error; err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	if arrivedToken.Token != candidateToken.Token {
		ctx.JSON(400, gin.H{
			"status": "invalid token",
		})
		return
	}

	if err := c.DB.
		Where("id = ?", &candidateToken.ID).
		Find(&candidateUser).Error; err != nil {
		fmt.Println("token is invalid")
		utils.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(200, &candidateUser)

}

func (c *AuthController) ChangeData(ctx *gin.Context) {
	var dbUser models.User
	var arrivedUser models.User

	if err := ctx.ShouldBindJSON(&arrivedUser); err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	if err := c.DB.
		Where("id = ?", &arrivedUser.ID).
		First(&dbUser).Error; err != nil {

		utils.ErrorResponse(ctx, err)
		return
	}

	if dbUser.Password != arrivedUser.Password {
		log.Fatal("password is invalid")
	}

	if err := c.DB.
		Model(&models.User{}).
		Where("id = ?", &dbUser.ID).
		Updates(&arrivedUser).Error; err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	type UserToken struct {
		Token string `json:"token"`
	}

	var userToken UserToken
	if err := ctx.ShouldBindJSON(&userToken); err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	err := c.DB.Unscoped().Delete(&models.Token{}, "token = ?", &userToken.Token).Error

	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": "success",
	})
}
