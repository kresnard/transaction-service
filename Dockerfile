FROM golang:1.22-alpine AS builder

# Install dependencies
RUN apk update && apk add --no-cache ca-certificates git librdkafka-dev musl-dev openssh-client && update-ca-certificates

# Set working dir
WORKDIR $GOPATH/src/app
COPY . .

# Fetch dependencies
RUN --mount=type=ssh go mod download
RUN --mount=type=ssh go mod tidy

# CMD go build -v
RUN GO111MODULE= on CGO_ENABLED=1 GOOS=linux go build -tags dynamic -o transaction-service cmd/main.go



#### small binary
FROM alpine:latest

# Install runtime dependencies
RUN apk update --no-cache && apk add --no-cache busybox-extras bash librdkafka-dev musl-dev tzdata
ENV TZ=Asia/Jakarta

# Create a non-root user and group
RUN addgroup -S trx-admin && adduser -S trx-admin -G trx-admin

# COPY the executeable & go.mod
COPY --from=builder /go/src/app/transaction-service /go/src/app/kpm/transaction-service
COPY --from=builder /go/src/app/go.mod /go/src/app/go.mod

# Set working directory
WORKDIR /go/src/app

# Ensure proper ownership of app files
RUN chown -R trx-admin:trx-admin /go/src/app

# Switch to the non-root user
USER trx-admin

# Expose the app port
EXPOSE 8080

ENTRYPOINT ["./transaction-service"]


