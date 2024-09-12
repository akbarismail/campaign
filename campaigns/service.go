package campaigns

type Service interface {
	FindCampaigns(userId int) ([]Campaigns, error)
}

type service struct {
	repo Repository
}

// FindCampaigns implements Service.
func (s *service) FindCampaigns(userId int) ([]Campaigns, error) {
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
