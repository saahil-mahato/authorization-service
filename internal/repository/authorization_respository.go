package repository

import (
	"github.com/saahil-mahato/authorization-service/internal/models"
	"gorm.io/gorm"
)

type AuthorizationRepository interface {
	Create(config *models.AuthorizationConfig) error
	GetByID(id uint) (*models.AuthorizationConfig, error)
	List() ([]*models.AuthorizationConfig, error)
	Update(config *models.AuthorizationConfig) error
	Delete(id uint) error
	FindByURLResourceRole(url, resource, role string) (*models.AuthorizationConfig, error)
}

type authorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizatonRepository(db *gorm.DB) AuthorizationRepository {
	return &authorizationRepository{db: db}
}

func (r *authorizationRepository) Create(config *models.AuthorizationConfig) error {
	return r.db.Create(config).Error
}

func (r *authorizationRepository) GetByID(id uint) (*models.AuthorizationConfig, error) {
	var config models.AuthorizationConfig
	err := r.db.First(&config, id).Error
	return &config, err
}

func (r *authorizationRepository) List() ([]*models.AuthorizationConfig, error) {
	var configs []*models.AuthorizationConfig
	err := r.db.Find(&configs).Error
	return configs, err
}

func (r *authorizationRepository) Update(config *models.AuthorizationConfig) error {
	return r.db.Save(config).Error
}

func (r *authorizationRepository) Delete(id uint) error {
	return r.db.Delete(&models.AuthorizationConfig{}, id).Error
}

func (r *authorizationRepository) FindByURLResourceRole(url, resource, role string) (*models.AuthorizationConfig, error) {
	var config models.AuthorizationConfig
	err := r.db.Where("url = ? AND resource = ? AND role = ?", url, resource, role).First(&config).Error
	return &config, err
}
