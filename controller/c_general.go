package controller

import (
	"backend-prokerin/db"
	"backend-prokerin/models"
	"backend-prokerin/services"
	"backend-prokerin/utilities"

	"github.com/gin-gonic/gin"
)

func ValidateUser(context *gin.Context) *models.UserAccount {
	db_connection := db.Connect()

	u_service := services.NewUserAccountService(db_connection)

	userId, errUserId := utilities.CurrentUser(context)
	if errUserId != nil {
		return nil
	}
	dataUser, errFindUser := u_service.FindUserById(userId)
	if errFindUser != nil {
		return nil
	}
	db.Close(db_connection)
	return &dataUser
}
