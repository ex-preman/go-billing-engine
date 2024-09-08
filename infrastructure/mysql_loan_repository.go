package infrastructure

import (
	"database/sql"

	"github.com/ex-preman/go-billing-engine/domain"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLLoanRepository struct {
	DB *sql.DB
}

// NewMySQLLoanRepository creates a new MySQL repository
func NewMySQLLoanRepository(dsn string) (*MySQLLoanRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLLoanRepository{DB: db}, nil
}

// Save stores the loan in the MySQL database
func (r *MySQLLoanRepository) Save(loan *domain.Loan) error {
	// Implement MySQL-specific save logic
	_, err := r.DB.Exec("INSERT INTO loans (id, outstanding, ...) VALUES (?, ?, ...)", loan.ID, loan.Outstanding)
	return err
}

// FindByID retrieves a loan from MySQL
func (r *MySQLLoanRepository) FindByID(id int) (*domain.Loan, error) {
	var loan domain.Loan
	err := r.DB.QueryRow("SELECT id, outstanding, ... FROM loans WHERE id = ?", id).Scan(&loan.ID, &loan.Outstanding)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
