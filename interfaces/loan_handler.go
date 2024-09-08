package interfaces

import (
	"net/http"
	"strconv"

	"github.com/ex-preman/go-billing-engine/application"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	LoanService *application.LoanService
}

// NewLoanHandler initializes LoanHandler
func NewLoanHandler(service *application.LoanService) *LoanHandler {
	return &LoanHandler{LoanService: service}
}

// CreateLoan handler to create a new loan
func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var req struct {
		ID           int     `json:"id"`
		Principal    float64 `json:"principal"`
		InterestRate float64 `json:"interest_rate"`
		Weeks        int     `json:"weeks"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.LoanService.CreateLoan(req.ID, req.Principal, req.InterestRate, req.Weeks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan created successfully"})
}

// GetOutstanding handler to fetch outstanding balance
func (h *LoanHandler) GetOutstanding(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	outstanding, err := h.LoanService.GetOutstanding(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"outstanding": outstanding})
}

// MakePayment handler to process loan payment
func (h *LoanHandler) MakePayment(c *gin.Context) {
	var req struct {
		Week   int     `json:"week"`
		Amount float64 `json:"amount"`
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.LoanService.MakePayment(id, req.Week, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment made successfully"})
}

// IsDelinquent handler to check if borrower is delinquent
func (h *LoanHandler) IsDelinquent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	isDelinquent, err := h.LoanService.IsDelinquent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_delinquent": isDelinquent})
}
