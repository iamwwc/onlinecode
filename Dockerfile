FROM debian:jessie as builder
ENV BUILD_DEPS "curl git"

COPY . /go/src/chaochaogege.com/onlinecode
WORKDIR /go/src/chaochaogege.com/onlinecode
RUN apt-get update && apt-get install curl git sudo --no-install-recommends -y
RUN curl -sL https://deb.nodesource.com/setup_8.x | sudo -E sh -;\
    apt-get install nodejs -y --no-install-recommends;

RUN cd ./client-side \
    && npm install && npm run build;
COPY ./sql ./client-side/dist/

RUN apt-get update && apt-get install ${BUILD_DEPS} -y --no-install-recommends;\
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
    dep ensure -update;\
    go install;

FROM golang:1.11.1-alpine3.7
RUN mkdir -p /go/src/chaochaogege.com/onlinecode
WORKDIR /go/src/chaochaogege.com/onlinecode
COPY --from=builder /go/bin/onlinecode .
COPY --from=builder /go/src/chaochaogege.com/onlinecode/client-side/dist/* .
EXPOSE 8086
ENTRYPOINT ["onlinecode"]