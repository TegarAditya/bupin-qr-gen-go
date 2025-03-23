FROM golang:1.24.1 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12 AS final

COPY --from=builder /go/bin/app /
COPY --from=builder /go/src/app/.env /
COPY --from=builder /go/src/app/assets /assets

CMD ["/app"]