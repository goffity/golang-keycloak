FROM golang:1.20-alpine3.18 AS builder

WORKDIR /identity-keycloak
COPY . .
RUN go test -v ./... && go install

FROM alpine:3.18 AS final
COPY --from=builder /go/bin/identity-keycloak /go/bin/identity-keycloak

EXPOSE 1323

CMD [ "/go/bin/identity-keycloak" ]