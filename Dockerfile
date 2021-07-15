FROM golang:1.16-alpine
WORKDIR /app
ADD go.mod go.sum *.go /app/
RUN go build -o /app/pgconn_check .

FROM alpine
WORKDIR /app
COPY --from=0 /app/pgconn_check /app/pgconn_check
ENTRYPOINT /app/pgconn_check
