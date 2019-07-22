FROM golang:1.12.7 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-play main.go
FROM scratch
COPY --from=builder /go/src/app/go-play /go-play
CMD ["./go-play"]