package controllers

import (
	"errors"
	"final_project/database"
	"final_project/helpers"
	"final_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//main function
func GetUsers(ctx *gin.Context) {
	var users = []models.User{}
	db := database.GetDB()
	db.Preload("User").Find((&users))
	ctx.JSON(http.StatusOK, users)
}

func RegisterUser(ctx *gin.Context) {
	NewUser := models.User{}
	db := database.GetDB()

	err := ctx.ShouldBindJSON(&NewUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Wrong Parameter")

		return
	}

	// _, errSearch := getUserbyName(NewUser.Username)
	// if errSearch == nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error_code": "REGUSR002",
	// 		"message":    "Username already exist. Please use another username.",
	// 	})

	// 	return
	// }

	errorCreate := db.Debug().Create(&NewUser).Error

	if errorCreate != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "REGUSR001",
			"message":    errorCreate.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"age":      NewUser.Age,
		"email":    NewUser.Email,
		"id":       NewUser.ID,
		"username": NewUser.Username,
	})
}

func LoginUser(ctx *gin.Context) {
	var new_account models.User
	err := ctx.ShouldBindJSON(&new_account)

	result, err := getUserbyName(new_account.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "USR001",
			"message":    "username not found. please check your username and try again.",
		})
		return
	}

	if !(helpers.CheckPasswordHash(new_account.Password, result.Password)) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "USR002",
			"message":    "Password does not match",
		})
		return
	}

	myToken := helpers.GenereteJwt(new_account.Username, new_account.Password)

	ctx.JSON(http.StatusOK, gin.H{
		"jwt": myToken,
	})
}

func UpdateUser(ctx *gin.Context) {
	user_id := ctx.Param("userId")
	var params_update models.User
	err := ctx.ShouldBindJSON(&params_update)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Wrong Parameter")

		return
	}

	result, err := getUserbyId(user_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "USR003",
			"message":    "user id not found. please check your user id and try again.",
		})
		return
	}

	result_update, err := updateAccount(result.ID, params_update.Email, params_update.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "USR004",
			"message":    "Failed to update account.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"email":      result_update.Email,
		"username":   result_update.Username,
		"age":        result.Age,
		"updated_at": result_update.UpdatedAt,
	})
}

//general function
func getUserbyId(user_id string) (models.User, error) {
	db := database.GetDB()

	user_account := models.User{}
	err := db.First(&user_account, "id = ?", user_id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_account, err
		}
		return user_account, err
	}

	return user_account, nil
}

func getUserbyName(username string) (models.User, error) {
	db := database.GetDB()

	user_account := models.User{}
	err := db.First(&user_account, "username = ?", username).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_account, err
		}
		return user_account, err
	}

	return user_account, nil
}

func updateAccount(user_id uint, email string, username string) (models.User, error) {
	db := database.GetDB()
	account_updated := models.User{}

	err := db.Model(&account_updated).Where("id = ?", user_id).Updates(
		models.User{
			Email:    email,
			Username: username,
		}).Error

	if err != nil {
		return account_updated, err
	}

	return account_updated, nil
}
