FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git
WORKDIR /src
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -o simple-uuid

FROM scratch
WORKDIR /bin
COPY --from=0 /src/simple-uuid .
CMD ["/bin/simple-uuid"]
