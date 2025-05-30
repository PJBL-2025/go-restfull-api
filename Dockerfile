# Gunakan base image resmi Go dengan versi terbaru (Go 1.24.1)
FROM golang:1.24.1

# Set working directory dalam container
WORKDIR /app

# Copy dependency file dan unduh semua modul
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy seluruh source code ke container
COPY . .

# Build binary Go
RUN go build -o main .

# Jalankan binary ketika container di-start
CMD ["./main"]