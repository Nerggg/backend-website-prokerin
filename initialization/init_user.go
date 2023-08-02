package initialization

import (
	"backend-prokerin/db"
	"backend-prokerin/models"
	"backend-prokerin/services"
)

func CreateSuperAdmin() {
	db_connection := db.Connect()
	u_service := services.NewUserAccountService(db_connection)
	_, err := u_service.FindUserByUsername("prokerin@gmail.com")
	if err != nil {
		user := models.UserAccount{
			Status:   1,
			Email:    "prokerin@gmail.com",
			NickName: "prokerin",
			Password: "prokerin123",
			Instansi: "Milestone 17",
			IsAdmin:  0,
		}
		_, err_create := u_service.Save(&user)
		if err_create != nil {
			panic("asd")
		}
	}
	db.Close(db_connection)
}
