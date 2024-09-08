package infrastructure

import (
	"log"

	"github.com/ex-preman/go-billing-engine/application"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Type  string
		MySQL struct {
			DSN string
		}
		PostgreSQL struct {
			DSN string
		}
	}
}

// DIContainer holds the application's dependencies
type DIContainer struct {
	Config      Config
	LoanRepo    application.LoanRepository
	LoanService *application.LoanService
}

// NewDIContainer initializes the DI container with configuration
func NewDIContainer() *DIContainer {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into config struct: %v", err)
	}

	// Choose the repository based on the database type
	var loanRepo application.LoanRepository
	switch config.Database.Type {
	case "inmemory":
		loanRepo = NewInMemoryLoanRepository()

	case "mysql":
		loanRepo, err = NewMySQLLoanRepository(config.Database.MySQL.DSN)
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}

	case "postgresql":
		loanRepo, err = NewPostgreSQLLoanRepository(config.Database.PostgreSQL.DSN)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}

	default:
		log.Fatalf("Unsupported database type: %s", config.Database.Type)
	}

	loanService := application.NewLoanService(loanRepo)

	return &DIContainer{
		Config:      config,
		LoanRepo:    loanRepo,
		LoanService: loanService,
	}
}
