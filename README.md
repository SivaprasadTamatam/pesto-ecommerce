## pesto-ecommerce

## Authentication Microservice README
# Overview
    This microservice provides user authentication functionalities for a simple e-commerce application. It manages user registration, login, and token-based authentication using JSON Web Tokens (JWT). The service is built with Go, utilizes SQLite for local data storage, and can be deployed using Docker for containerization.

# Features
- User registration with encrypted password storage
- User login with JWT token generation
- Token validation and authentication middleware
- SQLite database backend for data storage
- Docker support for easy deployment and scaling

# Prerequisites
Before running the authentication service, ensure you have the following installed:
- Docker: Install Docker
- Go (if building from source): Install Go

# Setup
# 1. Clone the Repository
- Clone the repository to your local machine:
- bash
- Copy code
- git clone https://github.com/your-username/auth-service.git
- cd auth-service

# 2. Environment Variables
- Create a .env file in the root directory (if not already provided) and set the following environment variables:
- plaintext
- Copy code
- JWT_SECRET=your_secret_key
- Adjust JWT_SECRET as needed for your application's security requirements.

# 3. Build and Run with Docker
- Build and run the Docker container:
    - bash
    - Copy code
    - docker-compose up --build
    - This command builds the Docker image and starts the authentication service container. The service should now be accessible at http://localhost:8080.

# 4. Testing
- Use Postman or curl commands to test the authentication endpoints:

    - Register User: POST http://localhost:8080/register
    - Login User: POST http://localhost:8080/login
    - Verify Token: GET http://localhost:8080/verify-token
# 5. Development (Optional)
 - If you prefer to run the service locally for development:
 - bash 
 - Copy code
 - Ensure SQLite is installed and accessible
    go mod download
    go run main.go

# 6. Accessing SQLite Database
    To access the SQLite database (test.db), you can use SQLite command-line tools or any SQLite database browser.
    
    API Endpoints
    POST /register: Register a new user with email and password.
    POST /login: Login with registered email and password to receive a JWT token.
    GET /verify-token: Verify JWT token for authentication.