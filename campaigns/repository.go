package campaigns

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindByUserId(userId int) ([]Campaigns, error)
}

type repository struct {
	db *gorm.DB
}

// FindAll implements Repository.
func (r *repository) FindAll() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Preload("CampaignImage", "campaign_image.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// FindByUserId implements Repository.
func (r *repository) FindByUserId(userId int) ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Where("id = ?", userId).Preload("CampaignImage", "campaign_image.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
