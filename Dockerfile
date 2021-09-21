FROM golang:1.16-alpine as builder

WORKDIR /go/src/bnj

# Get Reflex for live reload in dev
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy


COPY . .

RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/bnj/run .

EXPOSE 5200
CMD ["./run"]