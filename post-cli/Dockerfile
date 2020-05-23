FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git
RUN apk --no-cache add musl-dev
RUN apk add build-base gcc abuild binutils binutils-doc gcc-doc

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o colab-post-cli


FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD post.json /app/post.json
COPY --from=builder /app/colab-post-cli .

CMD ["./colab-post-cli"]
