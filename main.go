package main

import (
	"strconv"

	"github.com/ex-preman/go-billing-engine/infrastructure"
	"github.com/ex-preman/go-billing-engine/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the DI container
	container := infrastructure.NewDIContainer()

	// Create a new Gin router
	router := gin.Default()

	// Initialize the loan handler with the loan service
	loanHandler := interfaces.NewLoanHandler(container.LoanService)

	// Define API routes
	router.POST("/loan", loanHandler.CreateLoan)
	router.GET("/loan/:id/outstanding", loanHandler.GetOutstanding)
	router.POST("/loan/:id/payment", loanHandler.MakePayment)
	router.GET("/loan/:id/delinquent", loanHandler.IsDelinquent)

	// Run the server using the port from the config
	router.Run(":" + strconv.Itoa(container.Config.Server.Port))
}
