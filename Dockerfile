FROM golang:1.15-alpine AS builder
WORKDIR /app
COPY src/ /app
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch

LABEL maintainer leo@letnh.com

COPY --from=builder /bin/app /bin/app

EXPOSE 8080
CMD ["/bin/app"]
