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
	Save(input CreateTransactionInput) (Transaction, error)
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

func (s *service) Save(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignId
	transaction.UserId = input.User.ID
	transaction.Status = "Pending"

	newtransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newtransaction, err
	}

	return newtransaction, nil
}
