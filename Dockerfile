FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

#build go app
RUN go mod download
RUN go build -o Currency-operations ./cmd/currency-operations/main.go

CMD ["./Currency-operations"]