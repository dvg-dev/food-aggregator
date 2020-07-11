FROM golang:1.11 as builder

WORKDIR /food-aggregator

# Copy everything from the current directory to container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Start from scratch
FROM alpine:latest  

WORKDIR /go/src/github.com/dvg-dev/food-aggregator

#Copy the binary file from the previous stage
COPY --from=builder /food-aggregator/main .

COPY --from=builder /food-aggregator/urls.json .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
