# Golang Blockchain with React Frontend

A simplified blockchain implementation with proof-of-work mining, REST API microservice, and React frontend for visualization.

## Author
Rahman Azhar

## About
This project implements a basic blockchain with proof-of-work mining and provides both a REST API microservice and a React-based frontend interface. The blockchain maintains data integrity through SHA-256 hashing and includes validation mechanisms to ensure chain consistency.

## Features
- Blockchain implementation with proof-of-work mining
- Chain validation system
- REST API endpoints for blockchain interaction
- React frontend for blockchain visualization
- Docker containerization for easy deployment

## Project Structure
- `/pkg/blockchain`: Core blockchain implementation
- `/api`: REST API server implementation
- `/frontend`: React frontend application
- `main.go`: Backend application entry point
- `Dockerfile` & `docker-compose.yml`: Container configuration

## Running with Docker

1. Build and start the services:
```bash
docker-compose up --build
```

2. Access the application:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

## API Endpoints

### GET /blockchain
Returns the current state of the blockchain and its validity status.

Response format:
```json
{
    "blocks": [
        {
            "timestamp": 1234567890,
            "data": "Block Data",
            "hash": "block_hash",
            "prevHash": "previous_block_hash",
            "nonce": 123
        }
    ],
    "valid": true
}
```

### POST /block
Adds a new block to the blockchain.

Request format:
```json
{
    "data": "Your block data here"
}
```

## Technical Details
- Backend: Go
- Frontend: React with Chakra UI
- Containerization: Docker & Docker Compose
- Hashing: SHA-256
- State Management: React Hooks
- API Communication: Axios
- Styling: Chakra UI components

## Development

To run the services separately for development:

1. Backend:
```bash
go run main.go
```

2. Frontend:
```bash
cd frontend
npm install
npm start
```

## Features
- Real-time blockchain visualization
- Block mining with proof-of-work
- Chain validation checking
- Interactive block addition
- Responsive design
- Automatic chain updates
