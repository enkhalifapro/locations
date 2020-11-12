# builder image
FROM golang:1-alpine as gobuilder

RUN apk update
RUN apk add --no-cache bash curl git make

WORKDIR /locations

COPY . /locations

RUN make build

# actual image
FROM alpine

RUN apk update

COPY --from=gobuilder /locations/location /

EXPOSE 8000

ENTRYPOINT ["./location"]
