# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go env -w GO111MODULE=off && go build -o whereami main.go

# Production stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/whereami .
EXPOSE 8080
ENV PORT=8080
CMD ["./whereami"]
