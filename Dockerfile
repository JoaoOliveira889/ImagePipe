# Stage 1: Build the binary
FROM golang:1.25.5-alpine AS builder

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# IMPORTANT: We use TARGETARCH to make this work on any machine
ARG TARGETARCH
RUN CGO_ENABLED=1 GOOS=linux GOARCH=$TARGETARCH go build -o imagepipe .

# Stage 2: Final lightweight image
FROM alpine:latest
RUN apk add --no-cache libc6-compat

WORKDIR /data
COPY --from=builder /app/imagepipe /usr/local/bin/imagepipe

ENTRYPOINT ["imagepipe"]