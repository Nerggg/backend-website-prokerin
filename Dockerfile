FROM golang:1.19 as base


# ------------------------------------------------------------

FROM base as none
CMD ["echo", "Sorry, I am disabled"]

# ------------------------------------------------------------

FROM golang:alpine as builder

WORKDIR /go/src/backend-website-prokerin
# Get dependancies - will also be cached if we won't change mod/sum
COPY go.mod /go/src/backend-website-prokerin/
COPY go.sum /go/src/backend-website-prokerin/
RUN go mod download

COPY . /go/src/backend-website-prokerin
RUN go build -o ./dist/backend-website-prokerin

# ------------------------------------------------------------

FROM alpine:3.16.6 as local
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
  apk del tzdata

COPY ./config/config.yaml ./config.yaml
COPY --from=builder /go/src/backend-website-prokerin/dist/backend-website-prokerin .
EXPOSE 3000
ENTRYPOINT ["./backend-website-prokerin"]
