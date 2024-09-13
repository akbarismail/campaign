package campaigns

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindByUserId(userId int) ([]Campaigns, error)
	FindById(id int) (Campaigns, error)
	Save(campaign Campaigns) (Campaigns, error)
	Update(campaign Campaigns) (Campaigns, error)
}

type repository struct {
	db *gorm.DB
}

// Update implements Repository.
func (r *repository) Update(campaign Campaigns) (Campaigns, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Save implements Repository.
func (r *repository) Save(campaign Campaigns) (Campaigns, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// FindById implements Repository.
func (r *repository) FindById(id int) (Campaigns, error) {
	var campaign Campaigns
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// FindAll implements Repository.
func (r *repository) FindAll() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

// FindByUserId implements Repository.
func (r *repository) FindByUserId(userId int) ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
