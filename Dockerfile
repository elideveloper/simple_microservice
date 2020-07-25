FROM alpine:latest

RUN echo "installing go..." 
RUN apk add --no-cache git make musl-dev go

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

COPY . .

RUN go build -o ./my_app

CMD ["./my_app"]
EXPOSE 8080