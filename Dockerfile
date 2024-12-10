FROM golang:1.21 AS backend-builder
WORKDIR /app
COPY go.* ./
COPY pkg/ ./pkg/
COPY api/ ./api/
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o blockchain-server

FROM node:20 AS frontend-builder
WORKDIR /app
COPY frontend/ ./
RUN npm install
RUN npm run build

FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/blockchain-server ./
COPY --from=frontend-builder /app/build ./frontend/build
RUN apk add --no-cache ca-certificates

EXPOSE 8080
CMD ["./blockchain-server"]
