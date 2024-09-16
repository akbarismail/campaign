package transaction

import "gorm.io/gorm"

type Repository interface {
	FindCampaignId(campaignId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

// FindCampaignId implements Repository.
func (r *repository) FindCampaignId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
