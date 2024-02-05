FROM golang:1.21.6-alpine3.19 as build-stage
WORKDIR /workdir
COPY . . 
RUN go mod download && CGO_ENABLE=0 go build -ldflags "-s -w" -o app ./main.go

FROM alpine:3.19 as prod-stage
COPY --from=build-stage /workdir/app /

# grpc port
EXPOSE 9000 

# http port
EXPOSE 8080 

ENTRYPOINT [ "/app" ]
