package repositories

import (
	"errors"
	"go_starter/models"
	"go_starter/requests"
	"go_starter/trails"
	"gorm.io/gorm"
)

type PartnerRepository interface {
	//insert your function interface
	CreateSeller(model models.Partner) error
	FindByCredentials(request requests.LoginRequest) (*models.Partner, error)
	GetAllPartner() ([]models.Partner, error)
}

type partnerRepository struct{ db *gorm.DB }

func (p partnerRepository) GetAllPartner() ([]models.Partner, error) {
	var models []models.Partner
	err := p.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (p partnerRepository) CreateSeller(model models.Partner) error {
	hashPassword, err := trails.HashPassword(model.Password)
	if err != nil {
		return err
	}
	seller := models.Partner{
		Name:     model.Name,
		Username: model.Username,
		Password: hashPassword,
	}
	result := p.db.Create(&seller)
	if result.Error != nil {
		if result.Error.Error() == `ERROR: duplicate key value violates unique constraint "partners_username_key" (SQLSTATE 23505)` {
			return errors.New("username is exist")
		}
		return result.Error
	}
	return nil
}

func (p partnerRepository) FindByCredentials(request requests.LoginRequest) (*models.Partner, error) {
	var model models.Partner
	result := p.db.Where("username", request.Username).First(&model)
	if result.RowsAffected <= 0 {
		return nil, errors.New("partner not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	if !trails.CheckPasswordHash(request.Password, model.Password) {
		return nil, errors.New("wrong password")
	}
	if request.Username == model.Username && trails.CheckPasswordHash(request.Password, model.Password) {
		return &models.Partner{
			ID:       model.ID,
			Name:     model.Name,
			Username: model.Username,
		}, nil
	}
	return nil, errors.New("verify error")
}

func NewPartnerRepository(db *gorm.DB) PartnerRepository {
	// db.Migrator().DropTable(models.Partner{})
	//db.AutoMigrate(models.Partner{})
	return &partnerRepository{db: db}
}
