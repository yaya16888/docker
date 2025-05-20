# --- builder ---
    FROM golang:1.22-bullseye AS builder
    WORKDIR /src
    COPY ./go/go.mod ./go/go.sum ./
    RUN go mod download
    COPY ./go .
    RUN go install github.com/cosmtrek/air@latest && go test ./...
    
    # --- runner ---
    FROM golang:1.22-bullseye
    WORKDIR /app
    COPY --from=builder /src .
    COPY --from=builder /go/bin/air /usr/local/bin/air
    COPY .air.toml .
    EXPOSE 8080
    CMD ["air", "-c", ".air.toml"]