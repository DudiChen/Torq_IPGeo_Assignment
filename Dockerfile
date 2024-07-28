FROM golang:1.22-alpine AS build_base

# Set the Current Working Directory inside the container
WORKDIR /tmp/torq_ipgeo_app

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/torq_ipgeo_app .

# Start fresh from a smaller image
FROM alpine:3.9 

COPY --from=build_base /tmp/torq_ipgeo_app/out/torq_ipgeo_app /app/torq_ipgeo_app
COPY --from=build_base /tmp/torq_ipgeo_app/csv/data.csv /app/data.csv

ENV DATABASE_TYPE="csv" \
    DATABASE_PATH="/app/data.csv" \
    RATE_REQUESTS=1 \
    RATE_INTERVAL="1s"

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD ["/app/torq_ipgeo_app"]
