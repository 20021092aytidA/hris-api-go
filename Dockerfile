FROM golang:1.25.0

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /hris-api-go

EXPOSE 8080

CMD ["/hris-api-go"]