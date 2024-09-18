package transaction

import "gorm.io/gorm"

type Repository interface {
	FindCampaignId(campaignId int) ([]Transaction, error)
	FindUserId(userId int) ([]Transaction, error)
	FindById(id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

// FindById implements Repository.
func (r *repository) FindById(id int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ?", id).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Update implements Repository.
func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Save implements Repository.
func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

// FindUserId implements Repository.
func (r *repository) FindUserId(userId int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Where("user_id = ?", userId).Preload("Campaign.CampaignImages", "campaign_images.is_primary=1").Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
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
