# Goland-Jam

Goland-Jam is a sample server demonstrating how to set up a microservice with Go, integrate MongoDB, and generate API documentation using Swagger.

## Project Structure

```plaintext
Goland-Jam/
│
├── cmd/
│   └── main.go
│
├── pkg/
│   ├── controllers/
│   │   └── health.go
│   │   └── member.go
│   ├── routes/
│   │   └── routes.go
│   └── config/
│       └── config.go
│   └── models/
│       └── member.go
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum

```

## Installation and Setup
### 1. Install Dependencies
```
go mod tidy
```

### 2. Generate Swagger Documentation
```
swag init -g cmd/main.go
```

### 3. Update Docker Compose File
```
services:
  web:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - MONGO_URI=???
    depends_on:
      - mongo

  mongo:
    image: mongo:latest
    ports:
      - "27017:27017"
```

### 4. Build and Run Docker Image
```
docker-compose build
docker-compose up
```
### 5. Access Swagger UI
```
http://localhost:8080/swagger/index.html
```