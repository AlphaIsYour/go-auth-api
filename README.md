# 🔐 Go Auth API (Gin + JWT)

A simple authentication API built using Go and the Gin web framework. This project demonstrates how to handle user registration, login, and protected routes using JWT (JSON Web Token).

## ✨ Features

- ✅ User registration
- ✅ User login
- ✅ JWT-based authentication
- ✅ Protected route (`/profile`)
- ✅ Middleware for token validation
- ✅ In-memory user storage

## 🛠️ Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Authentication:** JWT (github.com/golang-jwt/jwt/v5)
- **Database:** In-memory (map) – easily extendable to SQLite

## 📦 Installation

```bash
# Clone this repo
git clone https://github.com/yourusername/go-auth-api.git
cd go-auth-api

# Initialize Go modules (if needed)
go mod init go-auth-api

# Download dependencies
go get

# Run the server
go run main.go
