FROM golang:1.17.0-alpine3.14 as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .


# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/svc

FROM alpine:3.14.1
# until https://github.com/alpinelinux/docker-alpine/issues/199#issuecomment-906122484 is resolved
RUN apk add --no-cache openssl-dev


COPY --from=build-env /go/bin/svc /go/bin/svc
ENTRYPOINT ["/go/bin/svc"]