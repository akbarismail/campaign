package campaigns

type Service interface {
	GetCampaigns(userId int) ([]Campaigns, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaigns, error)
}

type service struct {
	repo Repository
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
