FROM golang:latest

WORKDIR .

COPY . .

RUN go mod download 

RUN go build -o server .

EXPOSE 8075

CMD ["./server"]
