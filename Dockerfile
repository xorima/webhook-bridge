############################
# STEP 1 build executable binary
############################
FROM golang:1.23-alpine3.20 AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

# Fetch dependencies.
# Using go mod with go 1.11
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code
COPY main.go main.go
COPY cmd cmd
COPY docs docs
COPY internal internal

# Build the binary
RUN  go build -ldflags="-w -s" -o /go/bin/webhook-bridge

FROM scratch AS runner
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Create /etc/app/ directory
WORKDIR /etc/app/
# Copy our static executable
COPY --from=builder /go/bin/webhook-bridge /go/bin/webhook-bridge
EXPOSE 3000
USER appuser:appuser
ENTRYPOINT ["/go/bin/webhook-bridge"]
