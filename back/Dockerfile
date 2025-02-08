FROM golang:1.23 as builder 

WORKDIR /app  

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .

RUN go build -o sudoku-api ./api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sudoku-api .


EXPOSE 3004
CMD ["./sudoku-api"]
