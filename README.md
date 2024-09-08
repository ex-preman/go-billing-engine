# go-billing-engine

Billing Engine (System Design and abstraction)
We are building a billing system for our Loan Engine. Basically the job of a billing engine is to provide the
- Loan schedule for a given loan (when am i supposed to pay how much)
- Outstanding Amount for a given loan
- Status of weather the customer is Delinquent or not

We offer loans to our customers a 50 week loan for Rp 5,000,000/- , and a flat interest rate of 10% per annum.
This means that when we create a new loan for the customer (say loan id 100) then it needs to provide the billing schedule for the loan as 

W1: 110000

W2: 110000

W3: 110000

…

W50: 110000 

The Borrower repays the Amount every week. (assume that borrower can only pay the exact amount of payable that week or not pay at all) 

We need the ability to track the  Outstanding balance of the loan (defined as pending amount the borrower needs to pay at any point) eg. at the beginning of the week it is 5,500,000/- and it decreases as the borrower continues to make repayment, at the end of the loan it should be 0/-

Some customers may miss repayments, If they miss 2 continuous repayments they are delinquent borrowers.

To cover up for missed payments customers will need to make repayments for the remaining amounts. ie if there are 2 pending payments they need to make 2 repayments(of the exact amount).

We need to track the borrower if the borrower is Delinquent (any borrower that who’s not paid for last 2 repayments)

We are looking for at least the following methods to be implemented
- GetOutstanding: This returns the current outstanding on a loan, 0 if no outstanding(or closed),
- IsDelinquent: If there are more than 2 weeks of Non payment of the loan amount
- MakePayment: Make a payment of certain amount on the loan

\
\
Folder Structure

```markdown
go-loan-billing/
│
├── config/
│   └── config.yaml              # Configuration file for the application
│
├── domain/                      # Domain layer containing core business logic
│   └── loan.go                  # Loan entity with domain-specific logic
│
├── application/                 # Application layer containing service logic
│   └── loan_service.go          # LoanService for handling business operations
│
├── infrastructure/              # Infrastructure layer for dependencies like repositories
│   ├── loan_repository.go       # In-memory loan repository
│   └── di_container.go          # DI Container for setting up services, repositories, and config
│
├── interfaces/                  # Interfaces layer for exposing the API
│   └── loan_handler.go          # Gin HTTP handlers for loan operations
│
├── main.go                      # Main entry point of the application
│
├── go.mod                       # Go module file for dependencies
├── go.sum                       # Go checksum file for dependency versions
│
└── README.md                    # Project documentation
```

Example Configuration
InMemory: For testing or development without a database.

```yaml
server:
  port: 8080

database:
  type: "inmemory"
```
MySQL: For using MySQL as the database.
```yaml
server:
  port: 8080

database:
  type: "mysql"
  mysql:
    dsn: "user:password@tcp(localhost:3306)/dbname"
```
PostgreSQL: For using PostgreSQL as the database.
```yaml
server:
  port: 8080

database:
  type: "postgresql"
  postgresql:
    dsn: "postgres://user:password@localhost/dbname?sslmode=disable"
```
Update the dsn fields with your actual database credentials and connection information.

\
\
Access the API\
You can interact with the application through the following API endpoints:

Create a Loan

```http
POST /loan
```
Request body example:
```json
{
  "id": 1,
  "principal": 5000000,
  "interest_rate": 0.10,
  "weeks": 50
}
```

Get Outstanding Amount

```http
GET /loan/:id/outstanding
```

Make a Payment

```http
POST /loan/:id/payment
```
Request body example:
```json
{
  "week": 1,
  "amount": 110000
}
```

Check if Delinquent

```http
GET /loan/:id/delinquent
```
Replace :id with the actual loan ID for the requests.