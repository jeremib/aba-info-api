FROM golang:1.20-alpine3.17 as builder
COPY go.mod go.sum /go/src/egov/routing-number-info/
WORKDIR /go/src/egov/routing-number-info
RUN go mod download
COPY . /go/src/egov/routing-number-info
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/routing-number-info-api egov/routing-number-info

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/egov/routing-number-info/build/routing-number-info-api /usr/bin/routing-number-info-api
COPY data/banks.json /tmp/banks.json
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/routing-number-info-api", "/tmp/banks.json"]
