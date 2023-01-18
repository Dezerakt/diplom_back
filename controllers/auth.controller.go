package controllers

import (
	"diplom_back/models"
	"diplom_back/utils"
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

	token, err := bcrypt.GenerateFromPassword([]byte(signUpUser.Email), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
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
		"token": token,
	})
}

func (c *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user *models.User
	result := c.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if result := utils.VerifyPassword(user.Password, payload.Password); result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	token, err := bcrypt.GenerateFromPassword([]byte(payload.Email), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	var firstToken *models.Token
	c.DB.Find(&firstToken, "user_id = ?", user.ID)

	if firstToken == nil {
		err = c.DB.Create(&models.Token{
			Model:  gorm.Model{},
			UserID: user.ID,
			Token:  string(token),
		}).Error
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) GetInfoByToken(ctx *gin.Context) {
	var (
		candidateToken models.Token
		candidateUser  models.User
	)

	arrivedToken := ctx.Request.Header["Token"][0]
	if err := c.DB.Where("token = ?", &arrivedToken).Find(&candidateToken).Error; err != nil {
		log.Fatal(err.Error())
	}

	if arrivedToken != candidateToken.Token {
		ctx.JSON(400, gin.H{
			"status": "invalid token",
		})
	} else {
		if err := c.DB.
			Where("id = ?", &candidateToken.ID).
			Find(&candidateUser).Error; err != nil {

			log.Fatal(err.Error())
		}

		ctx.JSON(200, gin.H{
			"user": candidateUser,
		})
	}

}

func (c *AuthController) ChangeData(ctx *gin.Context) {
	var candidateToken models.Token
	var dbUser models.User

	arrivedToken := ctx.Request.Header["Token"][0]
	var arrivedUser models.User

	if err := ctx.ShouldBindJSON(&arrivedUser); err != nil {
		log.Fatal()
	}

	if err := c.DB.Where("token = ?", &arrivedToken).Find(&candidateToken).Error; err != nil {
		log.Fatal(err.Error())
	}

	if err := c.DB.
		Where("id = ?", &candidateToken.ID).
		Find(&dbUser).Error; err != nil {

		log.Fatal(err.Error())
	}

	if arrivedToken == candidateToken.Token {
		if err := c.DB.
			Model(&models.User{}).
			Where("id = ?", &dbUser.ID).
			Updates(&arrivedUser).Error; err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, gin.H{
			"status": "success",
		})
	} else {
		ctx.JSON(400, gin.H{
			"status": "token invalid",
		})
	}
}

func (c *AuthController) Logout(ctx *gin.Context) {
	type UserToken struct {
		Token string `json:"token"`
	}

	var userToken UserToken
	if err := ctx.ShouldBindJSON(&userToken); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err := c.DB.Where("token = ?", userToken.Token).Unscoped().Delete(&models.Token{}).Error

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"status": "success",
	})
}
