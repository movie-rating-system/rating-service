# Rating Service

The Rating Service manages user ratings and reviews for movies in the Movie Rating System. This microservice provides functionality for submitting, retrieving, and aggregating movie ratings.

## Architecture

This service is part of the [Movie Rating System](https://github.com/movie-rating-system) microservices architecture:

- **metadata-service** - Manages movie metadata
- **movie-service** - Aggregates data from metadata and rating services
- **rating-service** - This service (handles user ratings and reviews)
- **shared-components** - Common utilities and Docker configurations

## Features

- Submit and retrieve movie ratings
- Calculate aggregate ratings and statistics
- HTTP and gRPC APIs
- Protocol Buffer definitions for type safety
- In-memory storage (can be extended to use databases)
- Rating validation and business rules

## API

### gRPC API
Defined in `api/rating.proto`:
- `GetRatings(movieId)` - Get all ratings for a movie
- `PutRating(rating)` - Submit a new rating
- `GetAggregatedRating(movieId)` - Get aggregated rating statistics

### HTTP API
- `GET /rating/{movieId}` - Get aggregated rating for a movie
- `POST /rating` - Submit a new rating
- `GET /rating/{movieId}/all` - Get all individual ratings

## Getting Started

### Prerequisites
- Go 1.19 or higher
- Protocol Buffers compiler (protoc)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/movie-rating-system/rating-service.git
   cd rating-service
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate Protocol Buffer files (if needed):
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
          --go-grpc_out=. --go-grpc_opt=paths=source_relative \
          api/rating.proto
   ```

### Running the Service

```bash
# Run the service
go run cmd/main.go

# Or build and run
go build -o rating-service cmd/main.go
./rating-service
```

The service will start on:
- HTTP API: `http://localhost:8083`
- gRPC API: `localhost:8085`

### Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...
```

### Building

```bash
# Build binary
go build -o rating-service cmd/main.go

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o rating-service-linux cmd/main.go
```

## Docker

```bash
# Build Docker image
docker build -t rating-service .

# Run with Docker
docker run -p 8083:8083 -p 8085:8085 rating-service
```

## Configuration

The service can be configured through environment variables:
- `HTTP_PORT` - HTTP server port (default: 8083)
- `GRPC_PORT` - gRPC server port (default: 8085)

## API Examples

### Submit a Rating
```bash
curl -X POST http://localhost:8083/rating \
  -H "Content-Type: application/json" \
  -d '{
    "movieId": "1",
    "userId": "user123",
    "rating": 5,
    "review": "Excellent movie!"
  }'
```

### Get Aggregated Rating
```bash
curl http://localhost:8083/rating/1

# Response
{
  "movieId": "1",
  "averageRating": 4.2,
  "totalRatings": 150
}
```

### Get All Ratings for Movie
```bash
curl http://localhost:8083/rating/1/all
```

## Data Model

### Rating
```go
type Rating struct {
    MovieID string
    UserID  string
    Rating  int32  // 1-5 scale
    Review  string // Optional review text
}
```

### Aggregated Rating
```go
type AggregatedRating struct {
    MovieID       string
    AverageRating float64
    TotalRatings  int64
}
```

## Business Rules

- Ratings must be between 1 and 5
- One rating per user per movie (updates existing rating)
- Reviews are optional
- Aggregate ratings calculated in real-time

## Project Structure

```
├── api/                    # Protocol Buffer definitions
│   └── rating.proto
├── cmd/                    # Application entry points
│   ├── main.go
│   └── app.go
├── gen/                    # Generated Protocol Buffer code
├── internal/               # Internal application code
│   ├── config/            # Configuration management
│   ├── controller/        # Business logic controllers
│   ├── reporitory/        # Data access layer (note: typo in original)
│   └── service/           # Service layer
├── model/                 # Data models
└── go.mod
```

## Performance Considerations

- In-memory storage for fast access
- Efficient aggregation calculations
- Concurrent rating submissions handling
- Rate limiting for abuse prevention

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## Related Services

- [Metadata Service](https://github.com/movie-rating-system/metadata-service)
- [Movie Service](https://github.com/movie-rating-system/movie-service)
- [Shared Components](https://github.com/movie-rating-system/shared-components)

## License

This project is licensed under the MIT License.
