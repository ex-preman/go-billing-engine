package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ex-preman/go-billing-engine/infrastructure"
	"github.com/ex-preman/go-billing-engine/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	container := infrastructure.NewDIContainer()
	router := gin.Default()
	loanHandler := interfaces.NewLoanHandler(container.LoanService)

	router.POST("/loan", loanHandler.CreateLoan)
	router.GET("/loan/:id/outstanding", loanHandler.GetOutstanding)
	router.POST("/loan/:id/payment", loanHandler.MakePayment)
	router.GET("/loan/:id/delinquent", loanHandler.IsDelinquent)

	return router
}

func TestCreateLoanIntegration(t *testing.T) {
	router := setupRouter()

	loan := map[string]interface{}{
		"id":            1,
		"principal":     5000000,
		"interest_rate": 0.10,
		"weeks":         50,
	}

	jsonLoan, _ := json.Marshal(loan)
	req, _ := http.NewRequest(http.MethodPost, "/loan", bytes.NewBuffer(jsonLoan))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	t.Log(recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetOutstandingIntegration(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/loan/1/outstanding", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	t.Log(recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestMakePaymentIntegration(t *testing.T) {
	router := setupRouter()

	payment := map[string]interface{}{
		"week":   1,
		"amount": 110000,
	}

	jsonPayment, _ := json.Marshal(payment)
	req, _ := http.NewRequest(http.MethodPost, "/loan/1/payment", bytes.NewBuffer(jsonPayment))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	t.Log(recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestIsDelinquentIntegration(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/loan/1/delinquent", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	t.Log(recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}
