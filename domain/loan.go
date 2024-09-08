package domain

import "fmt"

type Loan struct {
	ID           int
	Principal    float64
	InterestRate float64
	TotalAmount  float64
	WeeklyAmount float64
	Weeks        int
	Outstanding  float64
	Payments     []bool
}

// NewLoan creates a new loan
func NewLoan(id int, principal float64, interestRate float64, weeks int) *Loan {
	totalAmount := principal + (principal * interestRate)
	weeklyAmount := totalAmount / float64(weeks)
	return &Loan{
		ID:           id,
		Principal:    principal,
		InterestRate: interestRate,
		TotalAmount:  totalAmount,
		WeeklyAmount: weeklyAmount,
		Weeks:        weeks,
		Outstanding:  totalAmount,
		Payments:     make([]bool, weeks),
	}
}

// GetOutstanding returns the current outstanding balance
func (l *Loan) GetOutstanding() float64 {
	return l.Outstanding
}

// IsDelinquent checks if the borrower missed 2 consecutive payments
func (l *Loan) IsDelinquent() bool {
	consecutiveMissed := 0
	for _, paid := range l.Payments {
		if !paid {
			consecutiveMissed++
			if consecutiveMissed >= 2 {
				return true
			}
		} else {
			consecutiveMissed = 0
		}
	}
	return false
}

// MakePayment processes a payment
func (l *Loan) MakePayment(week int, amount float64) error {
	if week < 1 || week > l.Weeks {
		return fmt.Errorf("invalid week")
	}
	if amount != l.WeeklyAmount {
		return fmt.Errorf("payment must be exactly: %.2f", l.WeeklyAmount)
	}
	if l.Payments[week-1] {
		return fmt.Errorf("payment for week %d already made", week)
	}

	l.Payments[week-1] = true
	l.Outstanding -= amount
	return nil
}
