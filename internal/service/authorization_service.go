package service

import (
	"github.com/saahil-mahato/authorization-service/internal/models"
	"github.com/saahil-mahato/authorization-service/internal/repository"
	"github.com/saahil-mahato/authorization-service/pkg/utils"
)

type AuthorizationService interface {
	CreateConfig(config *models.AuthorizationConfig) error
	GetConfig(id uint) (*models.AuthorizationConfig, error)
	ListConfigs() ([]*models.AuthorizationConfig, error)
	UpdateConfig(config *models.AuthorizationConfig) error
	DeleteConfig(id uint) error
	CheckAuthorization(url, resource, role, action string) (bool, string, error)
}

type authorizationService struct {
	repo repository.AuthorizationRepository
}

func NewAuthorizationService(repo repository.AuthorizationRepository) AuthorizationService {
	return &authorizationService{repo: repo}
}

func (s *authorizationService) CreateConfig(config *models.AuthorizationConfig) error {
	return s.repo.Create(config)
}

func (s *authorizationService) GetConfig(id uint) (*models.AuthorizationConfig, error) {
	return s.repo.GetByID(id)
}

func (s *authorizationService) ListConfigs() ([]*models.AuthorizationConfig, error) {
	return s.repo.List()
}

func (s *authorizationService) UpdateConfig(config *models.AuthorizationConfig) error {
	return s.repo.Update(config)
}

func (s *authorizationService) DeleteConfig(id uint) error {
	return s.repo.Delete(id)
}

func (s *authorizationService) CheckAuthorization(url, resource, role, action string) (bool, string, error) {
	config, err := s.repo.FindByURLResourceRole(url, resource, role)
	if err != nil {
		return false, "", err
	}

	authorized := false
	switch action {
	case "create":
		authorized = config.CanCreate
	case "read":
		authorized = config.CanRead
	case "update":
		authorized = config.CanUpdate
	case "delete":
		authorized = config.CanDelete
	}

	prismaSchema := utils.GeneratePrismaSchema([]*models.AuthorizationConfig{config})
	return authorized, prismaSchema, nil
}
