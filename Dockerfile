From golang:1.11.1-alpine3.7
COPY . /go/src/chaochaogege.com/onlinecode
RUN apk add curl \
            git ;\
            curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
            cd /go/src/chaochaogege.com/onlinecode; \
            pwd; \
            ls -ila; \
            dep ensure -update; \
            apk del curl; \
            apk del git;
RUN cd /go/src/chaochaogege.com/onlinecode; \
            go install \
            && ls | grep -v index.html | grep -v static | grep -v templates | xargs rm -Rf; \
            rm -f /go/bin/dep; \
            rm -Rf /go/pkg; \
            ls -ila;

WORKDIR /go/src/chaochaogege.com/onlinecode
EXPOSE 8086
ENTRYPOINT ["onlinecode"]
