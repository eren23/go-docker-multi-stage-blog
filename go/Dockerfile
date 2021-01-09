FROM golang:1.15
WORKDIR /src
COPY go.sum go.mod ./

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .
ENTRYPOINT ["/bin/app"]

FROM alpine
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]

