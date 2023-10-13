package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
)

type PartnerService interface {
	//Insert your function interface
	CreateSellerService(model models.Partner) error
	FindByCredentialsService(request requests.LoginRequest) (*models.Partner, error)
	ListPartnerService() ([]models.Partner, error)
}

type partnerService struct {
	repositoryPartner repositories.PartnerRepository
}

func (p partnerService) ListPartnerService() ([]models.Partner, error) {
	response, err := p.repositoryPartner.GetAllPartner()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p partnerService) CreateSellerService(model models.Partner) error {
	err := p.repositoryPartner.CreateSeller(model)
	if err != nil {
		return err
	}
	return nil
}

func (p partnerService) FindByCredentialsService(request requests.LoginRequest) (*models.Partner, error) {
	response, err := p.repositoryPartner.FindByCredentials(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewPartnerService(
	repositoryPartner repositories.PartnerRepository,
	// repo
) PartnerService {
	return &partnerService{
		repositoryPartner: repositoryPartner,
		//repo
	}
}
