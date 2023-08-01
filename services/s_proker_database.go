package services

import (
	"backend-prokerin/models"

	"github.com/rs/zerolog/log"
)

func (s *ProkerService) InsertToDatabase(disease *models.Proker) (*models.Proker, error) {

	err := s.db.Create(&disease).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Insert to Database")
		return &models.Proker{}, err
	}
	return disease, nil
}

// kalo bagian find ada error check bagian jenis data id nya..
func (s *ProkerService) Find(id string) (*models.Proker, error) {
	var data models.Proker
	err := s.db.
		Where("id=? AND status!=?", id, 0).
		First(&data).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Proker{}, err
	}
	return &data, nil
}

func (s *ProkerService) FindAll() []models.Proker {
	var dataAll []models.Proker
	s.db.Where("status!=?", 0).Find(&dataAll) // tambahkan syarat tertentu jika perlu
	return dataAll
}

func (s *ProkerService) UpdateRecord(skdr *models.Proker, id string, names []string) (*models.Proker, error) {
	var existing_record models.Proker
	err := s.db.Where("id=?", id).First(&existing_record).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Proker{}, err
	}

	err_update := s.db.Model(&existing_record).Select(names).Updates(skdr).Error
	if err_update != nil {
		log.Error().Err(err_update).Msg("Error Update to Database")
		return &models.Proker{}, err_update
	}
	return &existing_record, nil
}

func (s *ProkerService) LikeProker(id string, like bool) (*models.Proker, error) {
	var existing_record models.Proker
	err := s.db.Where("id=?", id).First(&existing_record).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Proker{}, err
	}
	if like {
		err_update := s.db.Model(&existing_record).Update("like", existing_record.Like+1).Error
		if err_update != nil {
			log.Error().Err(err_update).Msg("Error Update to Database")
			return &models.Proker{}, err_update
		}
	} else {
		err_update := s.db.Model(&existing_record).Update("like", existing_record.Like-1).Error
		if err_update != nil {
			log.Error().Err(err_update).Msg("Error Update to Database")
			return &models.Proker{}, err_update
		}
	}
	return &existing_record, nil
}

func (s *ProkerService) DeleteRecord(existing_record models.Proker) error {

	err_delete := s.db.Model(existing_record).Update("status", 0).Error
	if err_delete != nil {
		log.Error().Err(err_delete).Msg("Error Delete to Database")
		return err_delete
	}
	return nil
}
