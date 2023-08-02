package controller

import (
	"backend-prokerin/db"
	"backend-prokerin/models"
	"backend-prokerin/schema"
	"backend-prokerin/services"
	"backend-prokerin/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RegisterUserAccount(context *gin.Context) {
	var input schema.RegisterUserBodyReq
	db_connection := db.Connect()
	u_service := services.NewUserAccountService(db_connection)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	validator := validator.New()
	if err := validator.Struct(input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	user := models.UserAccount{
		Status:   1,
		Email:    input.Email,
		NickName: input.NickName,
		Password: input.Password,
		Instansi: input.Instansi,
		IsAdmin:  1,
	}

	_, err_create := u_service.Save(&user)
	if err_create != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_create})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}

func Login(context *gin.Context) {
	var input schema.LoginBodyReq

	db_connection := db.Connect()
	u_service := services.NewUserAccountService(db_connection)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "adasdasd"})
		return
	}

	validator := validator.New()
	if err := validator.Struct(input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "aaa"})
		return
	}

	user, err_find := u_service.FindUserByUsername(input.Username)
	if err_find != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "wrong username or password"})
		return
	}

	if utilities.ValidatePassword(user.Password, input.Password) {
		jwt, err_jwt := utilities.GenerateJWT(user)

		if err_jwt != nil {
			context.JSON(http.StatusOK, gin.H{"success": false, "error": err_jwt})
			return
		}
		data := schema.SuccesLogin{
			AccessToken: jwt,
		}

		context.JSON(http.StatusOK, gin.H{"data": data, "success": true})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": false, "message": "wrong username or password"})
	db.Close(db_connection)
}
