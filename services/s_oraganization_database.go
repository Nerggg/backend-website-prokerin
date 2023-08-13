package services

import (
	"backend-prokerin/models"

	"github.com/rs/zerolog/log"
)

func (s *OrganizationService) InsertToDatabase(disease *models.Organization) (*models.Organization, error) {

	err := s.db.Create(&disease).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Insert to Database")
		return &models.Organization{}, err
	}
	return disease, nil
}

// kalo bagian find ada error check bagian jenis data id nya..
func (s *OrganizationService) Find(id string) (*models.Organization, error) {
	var data models.Organization
	err := s.db.
		Where("id=? AND status!=?", id, 0).
		First(&data).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Organization{}, err
	}
	return &data, nil
}

func (s *OrganizationService) FindAll() []models.Organization {
	var dataAll []models.Organization
	s.db.Where("status!=?", 0).Find(&dataAll) // tambahkan syarat tertentu jika perlu
	return dataAll
}

func (s *OrganizationService) UpdateRecord(skdr *models.Organization, id string, names []string) (*models.Organization, error) {
	var existing_record models.Organization
	err := s.db.Where("id=?", id).First(&existing_record).Error
	if err != nil {
		log.Error().Err(err).Msg("Error Find to Database")
		return &models.Organization{}, err
	}

	err_update := s.db.Model(&existing_record).Select(names).Updates(skdr).Error
	if err_update != nil {
		log.Error().Err(err_update).Msg("Error Update to Database")
		return &models.Organization{}, err_update
	}
	return &existing_record, nil
}

func (s *OrganizationService) DeleteRecord(existing_record models.Organization) error {

	err_delete := s.db.Model(existing_record).Update("status", 0).Error
	if err_delete != nil {
		log.Error().Err(err_delete).Msg("Error Delete to Database")
		return err_delete
	}
	return nil
}
