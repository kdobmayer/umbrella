FROM golang:1.10
WORKDIR /go/src/github.com/kdobmayer/backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install ./...

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/bin/backend /bin/
CMD ["/bin/backend"]
EXPOSE 8080
