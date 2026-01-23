<!DOCTYPE html>
<html lang="en">

<body>

<div class="container py-5">

  <h1 class="mb-3">🏦 Go Bank Simulator</h1>
  <p class="lead">
    A full-featured <b>Banking Simulation API</b> built with <b>Go (Golang)</b>.
  </p>

  <hr>

  <h2>🚀 Features</h2>
  <ul>
    <li>Customer management</li>
    <li>Account creation & management</li>
    <li>Deposit & Withdraw</li>
    <li>Account to Account Transfer</li>
    <li>Customer to Customer Transfer</li>
    <li>Transaction history</li>
    <li>SQLite + GORM</li>
    <li>RESTful API</li>
    <li>Frontend integration</li>
  </ul>

  <h2>🧱 Tech Stack</h2>
  <ul>
    <li><b>Backend:</b> Go (Golang)</li>
    <li><b>Framework:</b> Gin</li>
    <li><b>ORM:</b> GORM</li>
    <li><b>Database:</b> SQLite</li>
    <li><b>Frontend:</b> HTML + Bootstrap + JS</li>
  </ul>

  <h2>📁 Project Structure</h2>
  <pre>
go-bank-simulator/
│
├── database/
├── handlers/
├── services/
├── repositorys/
├── models/
├── static/
├── templates/
├── main.go
└── README.md
  </pre>

  <h2>⚙️ Installation</h2>
  <pre>
git clone https://github.com/beyzkar/go-bank-simulator.git
cd go-bank-simulator
go mod tidy
go run main.go
  </pre>

  <h2>🌐 Run</h2>
  <pre>http://localhost:8080</pre>

  <h2>📌 API Endpoints</h2>

  <h4>Customers</h4>
  <ul>
    <li>POST /customers</li>
    <li>GET /customers</li>
    <li>GET /customers/:id</li>
    <li>DELETE /customers/:id</li>
  </ul>

  <h4>Accounts</h4>
  <ul>
    <li>POST /accounts</li>
    <li>GET /accounts/:id</li>
    <li>GET /customers/:id/accounts</li>
  </ul>

  <h4>Transactions</h4>
  <ul>
    <li>POST /accounts/:id/deposit</li>
    <li>POST /accounts/:id/withdraw</li>
    <li>GET /accounts/:id/transactions</li>
  </ul>

  <h4>Transfers</h4>
  <pre>
POST /accounts/transfer
{
  "fromAccountId": 1,
  "toAccountId": 2,
  "amount": 250
}
  </pre>

  <pre>
POST /transfer/by-customer
{
  "fromCustomerId": 1,
  "toCustomerId": 2,
  "amount": 250
}
  </pre>

  <h2>🧠 Highlights</h2>
  <ul>
    <li>Transaction-safe money transfers</li>
    <li>Automatic balance validation</li>
    <li>Clean layered architecture</li>
    <li>Multiple accounts per customer</li>
    <li>Production-ready structure</li>
  </ul>

  <h2>📌 Future Improvements</h2>
  <ul>
    <li>JWT Authentication</li>
    <li>Role-based access</li>
    <li>Transaction reports</li>
    <li>Docker support</li>
    <li>Pagination & filtering</li>
  </ul>
  
  <p>
    ⭐ If you like this project, give it a star on GitHub!
  </p>

</div>

</body>
</html>
