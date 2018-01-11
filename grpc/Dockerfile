FROM golang:1.9.2 as build_env

RUN go get -u google.golang.org/grpc

COPY . /go/src/github.com/bigflood/gostudy/grpc/
WORKDIR /go/src/github.com/bigflood/gostudy/grpc/
#RUN go build -o /ping_svr cmd/svr/main.go
#RUN CGO_ENABLED=0 go build -o /server cmd/svr/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o /server cmd/svr/main.go

FROM scratch
COPY --from=build_env /server /
ENTRYPOINT ["/server"]