package infrastructure

import (
	"errors"

	"github.com/ex-preman/go-billing-engine/domain"
)

// InMemoryLoanRepository stores loans in memory
type InMemoryLoanRepository struct {
	Loans map[int]*domain.Loan
}

// NewInMemoryLoanRepository creates an in-memory repo
func NewInMemoryLoanRepository() *InMemoryLoanRepository {
	return &InMemoryLoanRepository{
		Loans: make(map[int]*domain.Loan),
	}
}

// Save stores the loan in memory
func (r *InMemoryLoanRepository) Save(loan *domain.Loan) error {
	r.Loans[loan.ID] = loan
	return nil
}

// FindByID retrieves a loan by its ID
func (r *InMemoryLoanRepository) FindByID(id int) (*domain.Loan, error) {
	loan, exists := r.Loans[id]
	if !exists {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}
