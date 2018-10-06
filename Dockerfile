FROM debian:jessie-20180831 as builder
ENV BUILD_DEPS "curl git"
ENV CGO_ENABLED=0

# install golang
ENV GOLANG_VERSION 1.11.1
RUN apt-get update && apt-get install wget git ca-certificates -y --no-install-recommends;\
    wget -O go.tgz https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz > /dev/null;\
    tar -C /usr/local -xzf go.tgz > /dev/null;\
    rm go.tgz; \
    export PATH="/usr/local/go/bin:%PATH";\
    go version

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"



# install nodejs npm

RUN apt-get update && apt-get install curl git --no-install-recommends -y
RUN curl -sL https://deb.nodesource.com/setup_8.x |sh -;\
    apt-get install nodejs -y --no-install-recommends;


# build npm package
COPY . /go/src/chaochaogege.com/onlinecode
WORKDIR /go/src/chaochaogege.com/onlinecode
RUN cd ./client-side \
    && npm install \
    && npm run build;
RUN cp sql ./client-side/sql -R


# install dep and onlinecode dependencies
RUN apt-get update && apt-get install ${BUILD_DEPS} -y --no-install-recommends;\
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
    dep ensure -update;\
    go install;

# copy built package to final image
FROM golang:1.11.1-alpine3.7
RUN mkdir -p /go/src/chaochaogege.com/onlinecode
WORKDIR /go/src/chaochaogege.com/onlinecode
COPY --from=builder /go/bin/onlinecode ./
COPY --from=builder /go/src/chaochaogege.com/onlinecode/client-side/dist/* ./
EXPOSE 8086
ENTRYPOINT ["/go/src/chaochaogege.com/onlinecode"]
