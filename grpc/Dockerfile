FROM golang:1.9.2 as build_env

RUN git config --global http.sslVerify false
RUN go get -u -insecure google.golang.org/grpc

COPY cmd/svr /go/src/github.com/bigflood/gostudy/grpc/cmd/svr
COPY pb /go/src/github.com/bigflood/gostudy/grpc/pb
WORKDIR /go/src/github.com/bigflood/gostudy/grpc/
#RUN go build -o /ping_svr cmd/svr/main.go
#RUN CGO_ENABLED=0 go build -o /server cmd/svr/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o /server cmd/svr/main.go

FROM scratch
COPY --from=build_env /server /
ENTRYPOINT ["/server"]