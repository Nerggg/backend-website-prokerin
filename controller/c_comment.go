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

func AddComment(context *gin.Context) {
	var input schema.CommentBodyReq
	db_connection := db.Connect()
	proker := context.Params.ByName("proker")

	p_service := services.NewCommentService(db_connection)

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

	comment := models.Comment{
		Status:      1,
		IdCreator:   user.ID,
		IdProker:    proker,
		Description: input.Description,
		// Image: , ada cara khusus
		// TimeLineImage: ,
		Like: 0,
	}

	_, err_created := p_service.InsertToDatabase(&comment)
	if err_created != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_created})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}

func GetComment(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewCommentService(db_connection)
	proker := context.Params.ByName("proker")

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	comment, err_find := p_service.FindByProker(proker)
	if err_find != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_find})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": comment, "success": true})
	db.Close(db_connection)
}

func GetAllComment(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewCommentService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	all_comment := p_service.FindAll()

	context.JSON(http.StatusOK, gin.H{"data": all_comment, "success": true})
	db.Close(db_connection)
}

func LikeComment(context *gin.Context) {
	db_connection := db.Connect()
	id := context.Params.ByName("id")
	p_service := services.NewCommentService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	// like comment
	comment, err_like := p_service.LikeComment(id, true)
	if err_like != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_like})
		return
	}

	// update like comment in comment
	liked_comment := make(map[string]bool)
	if comment.Liked != nil {
		err_unmarshal_c := json.Unmarshal(comment.Liked, &liked_comment)
		if err_unmarshal_c != nil {
			context.JSON(http.StatusOK, gin.H{"success": false, "message": err_unmarshal_c})
			return
		}
	}

	liked_comment[user.ID] = true
	json_data_c, err_marhsal := json.Marshal(liked_comment)
	if err_marhsal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_marhsal})
		return
	}
	comment.Liked = json_data_c

	var changesComment []string = []string{"Liked"}

	_, err_update_c := p_service.UpdateRecord(comment, comment.ID, changesComment)
	if err_update_c != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_update_c})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
}

func UnLikeComment(context *gin.Context) {
	db_connection := db.Connect()
	id := context.Params.ByName("id")
	p_service := services.NewCommentService(db_connection)

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	// like comment
	comment, err_like := p_service.LikeComment(id, false)
	if err_like != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_like})
		return
	}

	// update like comment in comment
	liked_comment := make(map[string]bool)
	if comment.Liked != nil {
		err_unmarshal_c := json.Unmarshal(comment.Liked, &liked_comment)
		if err_unmarshal_c != nil {
			context.JSON(http.StatusOK, gin.H{"success": false, "message": err_unmarshal_c})
			return
		}
	}
	liked_comment[user.ID] = false
	json_data_c, err_marhsal := json.Marshal(liked_comment)
	if err_marhsal != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_marhsal})
		return
	}
	comment.Liked = json_data_c

	var changesComment []string = []string{"Liked"}

	_, err_update_c := p_service.UpdateRecord(comment, comment.ID, changesComment)
	if err_update_c != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_update_c})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
}

func DeleteComment(context *gin.Context) {
	db_connection := db.Connect()
	p_service := services.NewCommentService(db_connection)
	id := context.Params.ByName("id")

	// validate user
	user := ValidateUser(context)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid user"})
		return
	}

	comment, err_find := p_service.Find(id)
	if err_find != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_find})
		return
	}

	err_delete := p_service.DeleteRecord(*comment)
	if err_delete != nil {
		context.JSON(http.StatusOK, gin.H{"success": false, "message": err_delete})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": nil, "success": true})
	db.Close(db_connection)
}
