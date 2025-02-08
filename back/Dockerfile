FROM golang:1.23 as builder 

WORKDIR /app  

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/sudoku-api ./api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sudoku-api .

RUN chmod +x /root/sudoku-api

EXPOSE 3004
CMD ["./sudoku-api"]
