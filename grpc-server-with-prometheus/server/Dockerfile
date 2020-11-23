FROM golang:1.15.5
EXPOSE 9092
EXPOSE 9093
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]
