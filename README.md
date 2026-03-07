# Go Project

This repository contains Go code and related resources.

## Getting Started

1. **Clone the repository:**

   ```sh
   git clone https://github.com/MuhammadTopu/Go-initial.git
   cd Go-initial
   ```

2. **Run the code:**

   ```sh
   go run main.go
   ```

3. **Build the project:**
   ```sh
   go build
   ```
4. **Users**
   - POST /api/users/register
   - POST /api/users/login
   - GET /health
   - GET /api/users/profile
   - PUT /api/users/profile

## Project Structure

- `main.go` - Entry point of the application
- `README.md` - Project documentation
- `Commands.txt` - List of commands used in the project

## Architecture

│          BACKEND (Go + Gin)             │
│  Handlers → Services → Repositories     │
|       DATABASE (PostgreSQL 15)          │ 

## Requirements

- Go 1.18 or higher

## License

Kamran Hossain Topu
