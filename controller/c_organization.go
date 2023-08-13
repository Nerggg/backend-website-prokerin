package controller

import (
	"backend-prokerin/db"
	"backend-prokerin/models"
	"backend-prokerin/schema"
	"backend-prokerin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddOrganization(context *gin.Context) {
	var input schema.OrganizationBodyReq
	db_connection := db.Connect()

	p_service := services.NewOrganizationService(db_connection)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	validator := validator.New()
	if err := validator.Struct(input); err != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err})
		return
	}

	// validasi User
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	proker := models.Organization{
		Status:      1,
		Name:        input.Name,
		Logo:        input.Logo,
		Description: input.Description,
	}

	_, err_created := p_service.InsertToDatabase(&proker)
	if err_created != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_created})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}

func GetOrganization(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewOrganizationService(db_connection)
	id := context.Params.ByName("id")

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	proker, err_find := p_service.Find(id)
	if err_find != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_find})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": proker, "success": true})
	db.Close(db_connection)
}

func GetAllOrganization(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewOrganizationService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	all_proker := p_service.FindAll()

	context.JSON(http.StatusOK, gin.H{"data": all_proker, "success": true})
	db.Close(db_connection)
}
