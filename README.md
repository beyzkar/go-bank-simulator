🏦 Go Bank Simulator

A full-featured Banking Simulation API built with Go (Golang) using Gin, GORM, and SQLite.
This project simulates real-world banking operations such as customers, accounts, deposits, withdrawals, and transfers with transactional safety.

🚀 Features

✅ Customer management
✅ Account management
✅ Deposit & Withdraw operations
✅ Account-to-account transfer
✅ Customer-to-customer transfer
✅ Transaction history
✅ Safe database transactions (ACID)
✅ RESTful API
✅ Frontend integration (HTML + JS)
✅ SQLite database
✅ Clean architecture (Handler → Service → Repository)

🧱 Tech Stack
Layer	Technology
Language	Go (Golang)
Framework	Gin
ORM	GORM
Database	SQLite
Frontend	HTML + Bootstrap + Vanilla JS
Architecture	Layered (Handler / Service / Repository)
📁 Project Structure
go-bank-simulator/
│
├── database/
│   └── db.go
│
├── handlers/
│   ├── account_handler.go
│   ├── customer_handler.go
│   ├── transaction_handler.go
│
├── services/
│   ├── account_service.go
│   ├── transaction_service.go
│
├── repositorys/
│   ├── account_repository.go
│   ├── customer_repository.go
│   ├── transaction_repository.go
│
├── models/
│   ├── account.go
│   ├── customer.go
│   ├── transaction.go
│
├── static/
│   └── app.js
│
├── templates/
│   └── index.html
│
├── main.go
└── README.md

⚙️ Installation & Run
1️⃣ Clone the repository
git clone https://github.com/USERNAME/go-bank-simulator.git
cd go-bank-simulator

2️⃣ Install dependencies
go mod tidy

3️⃣ Run the project
go run main.go

4️⃣ Open in browser
http://localhost:8080

📌 API Endpoints
🧑 Customer
Method	Endpoint	Description
POST	/customers	Create customer
GET	/customers	List customers
GET	/customers/:id	Get customer
DELETE	/customers/:id	Delete customer
GET	/customers/search	Search customer
💳 Account
Method	Endpoint	Description
POST	/accounts	Create account
GET	/accounts/:id	Get account
GET	/customers/:id/accounts	Get customer accounts
DELETE	/accounts/:id	Delete account
💰 Transactions
Method	Endpoint	Description
POST	/accounts/:id/deposit	Deposit
POST	/accounts/:id/withdraw	Withdraw
GET	/accounts/:id/transactions	Transaction history
🔁 Transfers
Account → Account
POST /accounts/transfer

{
  "fromAccountId": 1,
  "toAccountId": 2,
  "amount": 250
}

Customer → Customer
POST /transfer/by-customer

{
  "fromCustomerId": 1,
  "toCustomerId": 2,
  "amount": 250
}


✔ Automatically selects the sender’s account with enough balance
✔ Uses database transactions
✔ Prevents invalid or unsafe transfers

🧠 Business Logic Highlights

✅ Transaction-safe money transfers

✅ Prevents negative balances

✅ Supports multiple accounts per customer

✅ Automatically selects valid sender account

✅ Uses DB-level atomic operations

✅ Clean separation of concerns

🖥️ Frontend Features

Customer search

Account listing

Deposit / Withdraw

Customer-to-customer transfer

Live balance update

Error & success feedback

🧪 Example Use Case

Create customers

Create accounts

Deposit money

Transfer between customers

View transaction history

📌 Future Improvements (Planned)

🔐 JWT Authentication

📊 Transaction history UI

🧾 PDF transaction export

💱 Multi-currency support

🧠 Fraud detection logic

🐳 Docker support

👨‍💻 Author

Beyza Karaalp
Backend Developer | Go Enthusiast

📌 GitHub: https://github.com/YOUR_USERNAME

⭐️ If you like this project

Give it a ⭐ on GitHub — it really helps!
