FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

#Download and install dependencies

RUN go get -d -v ./...

# Build the application

RUN go build -o backend .

# Expose the port that the application listens on

EXPOSE 8080
    
CMD ["./backend"]