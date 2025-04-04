FROM golang:1.24.1 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN chmod +x ./build.sh

RUN ./build.sh

FROM gcr.io/distroless/static-debian12 AS final

COPY --from=builder /go/src/app/build/bin/main /app
COPY --from=builder /go/src/app/assets /assets

CMD ["/app"]