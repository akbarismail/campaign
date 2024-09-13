package campaigns

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaigns, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
	UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaigns, error)
}

type service struct {
	repo Repository
}

// UpdateCampaign implements Service.
func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaigns, error) {
	campaign, err := s.repo.FindById(inputId.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.TotalAmount = inputData.TotalAmount
	campaign.Perks = inputData.Perks

	updateCampaign, err := s.repo.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}

	return updateCampaign, nil
}

// CreateCampaign implements Service.
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaigns, error) {
	campaigns := Campaigns{}
	campaigns.Name = input.Name
	campaigns.ShortDescription = input.ShortDescription
	campaigns.Description = input.Description
	campaigns.TotalAmount = input.TotalAmount
	campaigns.Perks = input.Perks
	campaigns.UserId = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	newSlug := slug.Make(slugCandidate)
	campaigns.Slug = newSlug

	newCampaign, err := s.repo.Save(campaigns)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

// GetCampaignById implements Service.
func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaigns, error) {
	campaign, err := s.repo.FindById(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// GetCampaigns implements Service.
func (s *service) GetCampaigns(userId int) ([]Campaigns, error) {
	if userId != 0 {
		campaigns, err := s.repo.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repo.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func NewService(repo Repository) Service {
	return &service{repo}
}
