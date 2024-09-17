package transaction

import (
	"campaign/campaigns"
	"errors"
)

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserId(userId int) ([]Transaction, error)
}

type service struct {
	repo               Repository
	campaignRepository campaigns.Repository
}

// GetTransactionsByUserId implements Service.
func (s *service) GetTransactionsByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repo.FindUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

// GetTransactionByCampaignId implements Service.
func (s *service) GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repo.FindCampaignId(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func NewService(repo Repository, campaignRepo campaigns.Repository) Service {
	return &service{repo, campaignRepo}
}
