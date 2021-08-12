FROM golang:1.16-alpine as builder

WORKDIR /app
COPY . .
# set CGO_ENABLED=0 to force an static build => smaller docker image
RUN CGO_ENABLED=0 go build .

FROM scratch
COPY --from=builder /app/ecommerce-service-starter /ecommerce-service

ENTRYPOINT ["/ecommerce-service"]