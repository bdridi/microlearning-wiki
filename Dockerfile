FROM golang
COPY . /go/src/wiki
RUN go install wiki
ENTRYPOINT /go/bin/wiki
EXPOSE 8081