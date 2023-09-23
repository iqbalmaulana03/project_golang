package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.CampaignRepository
}

type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.CampaignRepository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("Not an owner the campaign")
	}

	transactions, err := s.repository.GetByCampaignId(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transaction, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
