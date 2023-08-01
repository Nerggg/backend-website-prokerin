package services

import (
	"backend-prokerin/models"
	"errors"
	"html"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserAccountService) isEmailUnique(email string, isUpdate bool, id string) bool {
	var user_account models.UserAccount
	if isUpdate {
		err := s.db.Where("email=?", email).
			Where("id!=?", id).
			Where("status=?", 1).
			First(&user_account).Error
		return err == nil
	}
	err := s.db.
		Where("email=?", email).
		First(&user_account).Error
	return err == nil
}
func (s *UserAccountService) isNickNameUnique(nick_name string, isUpdate bool, id string) bool {
	var user_account models.UserAccount
	if isUpdate {
		err := s.db.Where("nick_name=?", nick_name).
			Where("id!=?", id).
			Where("status=?", 1).
			First(&user_account).Error
		return err == nil
	}
	err := s.db.
		Where("nick_name=?", nick_name).
		First(&user_account).Error
	return err == nil
}

func (s *UserAccountService) Save(data *models.UserAccount) (*models.UserAccount, error) {

	if s.isEmailUnique(data.Email, false, "") {
		return &models.UserAccount{}, errors.New("email not unique")
	}
	if s.isNickNameUnique(data.NickName, false, "") {
		return &models.UserAccount{}, errors.New("nick_name not unique")
	}

	err := s.db.Create(&data).Error
	if err != nil {
		return &models.UserAccount{}, err
	}
	return data, nil
}

func (s *UserAccountService) Update(data *models.UserAccount, names []string, id string) (*models.UserAccount, error) {
	var update_user models.UserAccount

	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		print(err)
	}
	data.Password = string(bytes)

	temp_email := html.EscapeString(strings.TrimSpace(data.Email))
	data.Email = temp_email

	// validate email unique
	if s.isEmailUnique(data.Email, true, id) {
		return &models.UserAccount{}, errors.New("email not unique")
	}
	if s.isNickNameUnique(data.NickName, true, id) {
		return &models.UserAccount{}, errors.New("nick_name not unique")
	}

	// do update
	err = s.db.Where("id=?", id).First(&update_user).Error
	if err != nil {
		return &models.UserAccount{}, err
	}

	err_update := s.db.
		Model(&update_user).
		Select(names).
		Updates(data).Error
	print("a")
	if err_update != nil {
		return &models.UserAccount{}, err_update
	}
	return data, nil
}

func (s *UserAccountService) FindUserById(id string) (models.UserAccount, error) {
	var user_account models.UserAccount
	err := s.db.
		Where("id=? AND status=?", id, 1).
		First(&user_account).Error
	if err != nil {
		return models.UserAccount{}, err
	}
	return user_account, nil
}

func (s *UserAccountService) FindUserByUsername(username string) (models.UserAccount, error) {
	var user_account models.UserAccount
	err := s.db.
		Where("(email=? OR nick_name=?) AND status=?", username, username, 1).
		First(&user_account).Error
	if err != nil {
		return models.UserAccount{}, err
	}
	return user_account, nil
}
func (s *UserAccountService) FindUserByToken(token string) (models.UserAccount, error) {
	var user_account models.UserAccount
	err := s.db.
		Where("token=? AND status=?", token, 1).
		First(&user_account).Error
	if err != nil {
		return models.UserAccount{}, err
	}
	return user_account, nil
}
func (s *UserAccountService) FindAll() []models.UserAccount {
	var user_account []models.UserAccount
	s.db.
		Where("status=?", 1).
		Find(&user_account)

	return user_account
}

func (s *UserAccountService) DeleteRecord(existing_record *models.UserAccount) error {

	err_delete := s.db.Model(existing_record).Update("status", 0).Error
	if err_delete != nil {
		log.Error().Err(err_delete).Msg("Error Delete to Database")
		return err_delete
	}
	return nil
}
