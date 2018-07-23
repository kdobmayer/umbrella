FROM golang:1.10
WORKDIR /go/src/github.com/kdobmayer/frontend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install ./...

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/bin/frontend /bin/
CMD ["/bin/frontend"]
EXPOSE 8081
