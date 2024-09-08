package application

import (
	"github.com/ex-preman/go-billing-engine/domain"
)

// LoanRepository defines the interface for loan repository
type LoanRepository interface {
	Save(loan *domain.Loan) error
	FindByID(id int) (*domain.Loan, error)
}

// LoanService contains business logic for managing loans
type LoanService struct {
	Repo LoanRepository
}

// NewLoanService initializes LoanService
func NewLoanService(repo LoanRepository) *LoanService {
	return &LoanService{Repo: repo}
}

// CreateLoan creates a new loan
func (s *LoanService) CreateLoan(id int, principal float64, interestRate float64, weeks int) error {
	loan := domain.NewLoan(id, principal, interestRate, weeks)
	return s.Repo.Save(loan)
}

// GetOutstanding fetches outstanding balance
func (s *LoanService) GetOutstanding(id int) (float64, error) {
	loan, err := s.Repo.FindByID(id)
	if err != nil {
		return 0, err
	}
	return loan.GetOutstanding(), nil
}

// MakePayment processes a payment
func (s *LoanService) MakePayment(id int, week int, amount float64) error {
	loan, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	err = loan.MakePayment(week, amount)
	if err != nil {
		return err
	}

	return s.Repo.Save(loan)
}

// IsDelinquent checks if the borrower is delinquent
func (s *LoanService) IsDelinquent(id int) (bool, error) {
	loan, err := s.Repo.FindByID(id)
	if err != nil {
		return false, err
	}
	return loan.IsDelinquent(), nil
}
