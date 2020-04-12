FROM golang:1.13 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

WORKDIR /opt/app
COPY ./tomozou /opt/app
RUN go build

# runtime image
FROM alpine
COPY --from=builder /opt/app /opt/app

CMD /opt/app/tomozou
