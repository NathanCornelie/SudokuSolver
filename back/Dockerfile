FROM golang:1.21 as builder 

WORKDIR /app  

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .

RUN go build -o sudoku-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .


EXPOSE 3004
CMD ["./sudoku-api"]
