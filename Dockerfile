FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app *.go

FROM scratch
COPY --from=build /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]