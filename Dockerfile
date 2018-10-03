From golang:1.10.4-alpine3.7
COPY . /go/src/chaochaogege.com/onlinecode
RUN apk add curl; \
             git; \
             curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
             cd /go/src/chaochaogege.com/onlinecode; \
             dep ensure -update -vendor-only &&  go install; \
             ls | grep -v index.html | grep -v static | grep -v templates | xargs rm -Rf; \
             rm -f /go/bin/dep; \
             rm -Rf /go/pkg; \
             apk del git; \
             apk del curl;
WORKDIR /go/src/chaochaogege.com/onlinecode
EXPOSE 8086
ENTRYPOINT ["onlinecode"]
