package campaign

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
}

type campaignService struct {
	repository CampaignRepository
}

func NewCampaignService(repository CampaignRepository) *campaignService {
	return &campaignService{repository}
}

func (s *campaignService) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *campaignService) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
