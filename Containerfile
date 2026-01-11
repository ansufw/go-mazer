# Stage 1: Build frontend assets
FROM oven/bun:1 AS frontend
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile
COPY . .
RUN bun run build

# Stage 2: Build Go binary
FROM golang:1.25-alpine AS backend
WORKDIR /app
# Install git for go mod download and build tools
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Install templ tool
RUN go install github.com/a-h/templ/cmd/templ@v0.3.960
# Generate templ code
RUN templ generate
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 3: Final image
FROM alpine:latest
WORKDIR /app
# Install minimal dependencies
RUN apk --no-cache add ca-certificates

# Copy binary and config
COPY --from=backend /app/main .
COPY --from=backend /app/config.yaml .

# Copy all assets (source + compiled) to match app.Static configuration
COPY --from=frontend /app/views/assets ./views/assets

EXPOSE 8080
CMD ["./main"]
