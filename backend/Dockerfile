FROM golang:1.25-alpine

WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN go build -o main main.go

# Expose port backend
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
