From golang:1.10.4-alpine3.7

COPY . /go/src/chaochaogege.com/onlinecode
RUN cd /go/src/chaochaogege.com/onlinecode && go install
WORKDIR /go/src/chaochaogege.com/onlinecode
EXPOSE 8086
RUN onlinecode
