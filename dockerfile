FROM golang:1.19-alpine

WORKDIR /go/src/bernie.top/dymest/loan-backend/

COPY . .

RUN go build -o app .

RUN echo "6"

EXPOSE 8080

CMD ["./app"]
