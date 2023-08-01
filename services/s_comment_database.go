package services

import (
	"backend-prokerin/models"

	"github.com/rs/zerolog/log"
)

func (s *CommentService) InsertToDatabase(disease *models.Comment) (*models.Comment, error) {

	err := s.db.Create(&disease).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Insert to Database")
		return &models.Comment{}, err
	}
	return disease, nil
}

// kalo bagian find ada error check bagian jenis data id nya..
func (s *CommentService) Find(id string) (*models.Comment, error) {
	var data models.Comment
	err := s.db.
		Where("id=? AND status!=?", id, 0).
		First(&data).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Comment{}, err
	}
	return &data, nil
}
func (s *CommentService) FindByProker(id string) ([]models.Comment, error) {
	var data []models.Comment
	err := s.db.
		Where("id_proker=? AND status!=?", id, 0).
		Find(&data).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return nil, err
	}
	return data, nil
}

func (s *CommentService) FindAll() []models.Comment {
	var dataAll []models.Comment
	s.db.Where("status!=?", 0).Find(&dataAll) // tambahkan syarat tertentu jika perlu
	return dataAll
}

func (s *CommentService) UpdateRecord(skdr *models.Comment, id string, names []string) (*models.Comment, error) {
	var existing_record models.Comment
	err := s.db.Where("id=?", id).First(&existing_record).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Comment{}, err
	}

	err_update := s.db.Model(&existing_record).Select(names).Updates(skdr).Error
	if err_update != nil {
		log.Error().Err(err_update).Msg("Error Update to Database")
		return &models.Comment{}, err_update
	}
	return &existing_record, nil
}

func (s *CommentService) LikeComment(id string, like bool) (*models.Comment, error) {
	var existing_record models.Comment
	err := s.db.Where("id=?", id).First(&existing_record).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Comment{}, err
	}
	if like {
		err_update := s.db.Model(&existing_record).Update("like", existing_record.Like+1).Error
		if err_update != nil {
			log.Error().Err(err_update).Msg("Error Update to Database")
			return &models.Comment{}, err_update
		}
	} else {
		err_update := s.db.Model(&existing_record).Update("like", existing_record.Like-1).Error
		if err_update != nil {
			log.Error().Err(err_update).Msg("Error Update to Database")
			return &models.Comment{}, err_update
		}
	}
	return &existing_record, nil
}

func (s *CommentService) DeleteRecord(existing_record models.Comment) error {

	err_delete := s.db.Model(existing_record).Update("status", 0).Error
	if err_delete != nil {
		log.Error().Err(err_delete).Msg("Error Delete to Database")
		return err_delete
	}
	return nil
}
