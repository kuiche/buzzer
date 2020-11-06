FROM golang:1.15 as builder
COPY ./ /gobuild/
WORKDIR /gobuild
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o buzzer
RUN ls

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /gobuild/buzzer ./buzzer
COPY --from=builder /gobuild/public/ ./public/
CMD ["./buzzer"]
