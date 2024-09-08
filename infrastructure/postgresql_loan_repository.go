package infrastructure

import (
	"database/sql"

	"github.com/ex-preman/go-billing-engine/domain"
	_ "github.com/lib/pq"
)

type PostgreSQLLoanRepository struct {
	DB *sql.DB
}

// NewPostgreSQLLoanRepository creates a new PostgreSQL repository
func NewPostgreSQLLoanRepository(dsn string) (*PostgreSQLLoanRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &PostgreSQLLoanRepository{DB: db}, nil
}

// Save stores the loan in the PostgreSQL database
func (r *PostgreSQLLoanRepository) Save(loan *domain.Loan) error {
	// Implement PostgreSQL-specific save logic
	_, err := r.DB.Exec("INSERT INTO loans (id, outstanding, ...) VALUES ($1, $2, ...)", loan.ID, loan.Outstanding)
	return err
}

// FindByID retrieves a loan from PostgreSQL
func (r *PostgreSQLLoanRepository) FindByID(id int) (*domain.Loan, error) {
	var loan domain.Loan
	err := r.DB.QueryRow("SELECT id, outstanding, ... FROM loans WHERE id = $1", id).Scan(&loan.ID, &loan.Outstanding)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
