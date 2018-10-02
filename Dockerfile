From golang:1.10.4-alpine3.7

RUN apk add curl \
             git && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . /go/src/chaochaogege.com/onlinecode
RUN cd /go/src/chaochaogege.com/onlinecode\
     && dep ensure -update \
     && go install \
     && ls | grep -v index.html | grep -v static | grep -v templates | xargs rm -Rf

WORKDIR /go/src/chaochaogege.com/onlinecode
EXPOSE 8086
ENTRYPOINT ["onlinecode"]
