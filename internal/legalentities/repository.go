package legalentities

import "github.com/krisch/crm-backend/pkg/postgres"

type LegalERepository interface {
	CreateBankAccount(message LegalEntities) (LegalEntities, error)
	GetAllBankAccounts() ([]LegalEntities, error)
	UpdateBankAccount(id int, message LegalEntities) (LegalEntities, error)
	DeleteBankAccount(id int) error
}

type legalentitiesRepository struct {
	gorm *postgres.GDB
}

func (r *legalentitiesRepository) CreateBankAccount(message LegalEntities) (LegalEntities, error) {
	result := r.gorm.DB.Create(&message)
	if result.Error != nil {
		return LegalEntities{}, result.Error
	}
	return message, nil
}

func (r *legalentitiesRepository) GetAllBankAccounts() ([]LegalEntities, error) {
	var messages []LegalEntities
	err := r.gorm.DB.Find(&messages).Error
	return messages, err
}

func (r *legalentitiesRepository) UpdateBankAccount(id int, message LegalEntities) (LegalEntities, error) {
	return message, r.gorm.DB.Model(&LegalEntities{}).Where("id = ?", id).Updates(message).Error
}

func (r *legalentitiesRepository) DeleteBankAccount(id int) error {
	return r.gorm.DB.Delete(&LegalEntities{}, id).Error
}

func NewRepository(gorm *postgres.GDB) LegalERepository {
	return &legalentitiesRepository{gorm: gorm}
}
