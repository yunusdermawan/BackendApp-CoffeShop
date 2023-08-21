# Create container
FROM golang:1.20.6-alpine

# Create folder
WORKDIR /goapp

# Copy all file
COPY . .

# Install dependencies and build
RUN go mod download
RUN go build -v -o /goapp/goback ./cmd/main.go

# Open port for app
EXPOSE 8001

# Run app
ENTRYPOINT [ "/goapp/goback" ]