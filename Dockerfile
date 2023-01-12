# Create build stage based on buster image
FROM golang:1.16-buster AS builder
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./
# Install any required modules
RUN go mod download
# Copy over Go source code
COPY *.go ./
# Run the Go build and output binary under go_simple_proxy
RUN go build -o /go_simple_proxy
# Create a new release build stage
FROM gcr.io/distroless/base-debian10
# Set the working directory to the root directory path
WORKDIR /
# Copy over the binary built from the previous stage
COPY --from=builder /go_simple_proxy /go_simple_proxy
# Run the app binary when we run the container
ENTRYPOINT ["/go_simple_proxy"]