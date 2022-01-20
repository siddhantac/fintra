package service

//go:generate moq -out service_mock_test.go . Repository

import (
	"fmt"
	"time"

	"github.com/siddhantac/fintra/domain"
)

const (
	dateLayout = "2006-01-02"
)

type Time string

func NewTime(t time.Time) Time {
	return Time(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}

type Service struct {
	repo Repository
}

type Repository interface {
	Insert(*domain.Transaction) error
	GetByID(string) (*domain.Transaction, error)
	GetAll() ([]*domain.Transaction, error)
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTransaction(id string) (*domain.Transaction, error) {
	fmt.Println(">> svc", id)
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTransactions() ([]*domain.Transaction, error) {
	return s.repo.GetAll()
}

func (s *Service) NewTransaction(amount float64, isDebit bool, date, category, transactionType, description, account string) (*domain.Transaction, error) {
	d, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	intAmt := int(amount * 100)

	transaction, err := domain.NewTransaction(intAmt, d, isDebit, category, transactionType, description, account)
	if err != nil {
		return nil, fmt.Errorf("domain.NewTransaction: %w", err)
	}

	if err := s.repo.Insert(transaction); err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}

	return transaction, nil
}
