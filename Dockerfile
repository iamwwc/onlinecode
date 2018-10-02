From golang:1.10.4-alpine3.7

COPY . /go/src/onlinecode
RUN cd /go/src/onlinecode && go install
WORKDIR /go/src/onlinecode
EXPOSE 8086
RUN onlinecode
