package controller

import (
	"backend-prokerin/db"
	"backend-prokerin/models"
	"backend-prokerin/schema"
	"backend-prokerin/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddProker(context *gin.Context) {
	var input schema.ProkerBodyReq
	db_connection := db.Connect()

	p_service := services.NewProkerService(db_connection)

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

	proker := models.Proker{
		Status:      1,
		IdCreator:   user.ID,
		Name:        input.Name,
		Description: input.Description,
		// Image: , ada cara khusus
		// TimeLineImage: ,
		Like: 0,
	}

	_, err_created := p_service.InsertToDatabase(&proker)
	if err_created != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_created})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}

func GetProker(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewProkerService(db_connection)
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

func GetAllProker(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewProkerService(db_connection)

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

func LikeProker(context *gin.Context) {
	db_connection := db.Connect()
	id := context.Params.ByName("id")
	p_service := services.NewProkerService(db_connection)
	u_service := services.NewUserAccountService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	// like proker
	_, err_like := p_service.LikeProker(id, true)
	if err_like != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_like})
		return
	}

	// update like proker in user account
	var liked_data map[string]bool
	err_unmarshal := json.Unmarshal(user.LikedProker, &liked_data)
	if err_unmarshal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_unmarshal})
		return
	}

	liked_data[id] = true

	json_data, err_marhsal := json.Marshal(liked_data)
	if err_marhsal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_marhsal})
		return
	}

	user.LikedProker = json_data
	var changesUser []string = []string{"LikedProker"}

	_, err_update := u_service.Update(user, changesUser, user.ID)
	if err_update != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_update})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
}
func UnLikeProker(context *gin.Context) {
	db_connection := db.Connect()
	id := context.Params.ByName("id")
	p_service := services.NewProkerService(db_connection)
	u_service := services.NewUserAccountService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	// like proker
	_, err_like := p_service.LikeProker(id, false)
	if err_like != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_like})
		return
	}

	// update like proker in user account
	var liked_data map[string]bool
	err_unmarshal := json.Unmarshal(user.LikedProker, &liked_data)
	if err_unmarshal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_unmarshal})
		return
	}

	liked_data[id] = false

	json_data, err_marhsal := json.Marshal(liked_data)
	if err_marhsal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_marhsal})
		return
	}

	user.LikedProker = json_data
	var changesUser []string = []string{"LikedProker"}

	_, err_update := u_service.Update(user, changesUser, user.ID)
	if err_update != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_update})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
}

func DeleteProker(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewProkerService(db_connection)
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

	err_delete := p_service.DeleteRecord(*proker)
	if err_delete != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_delete})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}
