FROM golang:1.22.1-alpine3.19
LABEL authors="ejimenezr21@gmail.com"

# Set working directory
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy rest of files
COPY . .

# Install additional development dependencies and Build application
RUN go install github.com/cosmtrek/air@v1.45.0 \
    # && go install github.com/go-delve/delve/cmd/dlv@1.22.1 \
    && go build -gcflags="-N -l" -o /bin/app cmd/main.go

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]
