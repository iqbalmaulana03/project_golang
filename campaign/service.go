package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(campaignId GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
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

func (s *campaignService) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.User.ID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *campaignService) UpdateCampaign(campaignId GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(campaignId.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != input.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	campaignNew, err := s.repository.Update(campaign)
	if err != nil {
		return campaignNew, err
	}

	return campaignNew, nil
}
