FROM golang:1.26.4 AS builder
WORKDIR /app
COPY . ./
RUN make xrpc
RUN CGO_ENABLED=0 GOOS=linux go build -v -o statusphere .

FROM gcr.io/distroless/static-debian13
COPY --from=builder /app/statusphere /usr/local/bin/statusphere
WORKDIR /data
ENTRYPOINT ["/usr/local/bin/statusphere"]
CMD ["statusphere"]
