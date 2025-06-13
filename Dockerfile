FROM golang:1.24.3 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o statusphere .

FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/statusphere /usr/local/bin/statusphere
RUN mkdir /data
WORKDIR /data
ENTRYPOINT ["/usr/local/bin/statusphere"]
CMD ["statusphere"]
