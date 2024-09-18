package transaction

import (
	"campaign/campaigns"
	"campaign/payment"
	"errors"
	"strconv"
)

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repo               Repository
	campaignRepository campaigns.Repository
	paymentService     payment.Service
}

// ProcessPayment implements Service.
func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderId)

	trx, err := s.repo.FindById(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		trx.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		trx.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		trx.Status = "cancelled"
	}

	updateTransaction, err := s.repo.Update(trx)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindById(updateTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid" {
		campaign.BackerCount += 1
		campaign.CurrentAmount += updateTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateTransaction implements Service.
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignId
	transaction.UserId = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repo.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = s.repo.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
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

func NewService(repo Repository, campaignRepo campaigns.Repository, paymentService payment.Service) Service {
	return &service{repo, campaignRepo, paymentService}
}
