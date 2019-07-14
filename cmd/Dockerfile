FROM golang:1.12-alpine as builder
COPY . /build
WORKDIR /build

ENV GO111MODULE=on

RUN apk add git coreutils make \
    && go mod vendor
RUN make bin

FROM scratch
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/bin/server ./

EXPOSE 8080
CMD ["/server", "-database-path=db/payment.store"]