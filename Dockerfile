# Builder
FROM golang AS builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/olivere/chrono
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

# Image
FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/olivere/chrono/chrono /chrono
CMD ["/chrono"]
