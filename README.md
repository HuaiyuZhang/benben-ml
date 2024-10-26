# benben-ml: Machine Learning Model Serving Platform

A high-performance service built in Go that serves ML models with RESTful API support.

## Features (Current)
- Simple HTTP server for model serving
- Basic model prediction endpoint
- Health check endpoint
- Configuration management
- Simple model interface

## Planned Features
- Model versioning
- A/B testing capabilities
- Real-time model performance monitoring
- Automatic model reloading
- Support for multiple ML framework formats
- Load balancing and request batching
- gRPC support

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Git

### Installation
1. Clone the repository:
```bash
git clone https://github.com/HuaiyuZhang/benben-ml.git
cd benben-ml
```

2. Install dependencies:
```bash
go mod download
```

3. Build the server:
```bash
go build -o benben-ml ./cmd/server
```

4. Run the server:
```bash
./benben-ml
```

### API Usage

#### Make a Prediction
```bash
curl -X POST http://localhost:8080/predict \
  -H "Content-Type: application/json" \
  -d '{"features": [1.0, 2.0, 3.0, 4.0]}'
```

#### Health Check
```bash
curl http://localhost:8080/health
```

## Project Structure
- `cmd/server`: Application entry point
- `internal/api`: HTTP handlers
- `internal/config`: Configuration management
- `internal/model`: Model interface and implementations
- `internal/server`: Server implementation
- `pkg/predictor`: Prediction interface
