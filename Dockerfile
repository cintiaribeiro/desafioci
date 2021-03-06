FROM golang:stretch AS builder 
ADD ./src/main.go /go/src/app/
RUN cd /go/src/app && go build -o goapp

FROM scratch
WORKDIR /go/src/app
COPY --from=builder /go/src/app/goapp /go/src/app
ENTRYPOINT ["/go/src/app/goapp"]