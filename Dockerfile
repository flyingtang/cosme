FROM golang:1.11-alpine as buildgolang


WORKDIR /go/src/cosme

COPY . .

RUN apk add git \
    && rootdir=`pwd` \
    && golangpath="$rootdir/vendor/golang.org/x"  \
    && mkdir -p $golangpath && ls -la $golangpath && cd $golangpath \
    && git clone https://github.com/golang/crypto.git && git clone https://github.com/golang/sys.git \
    && cd $rootdir \
    && go get -u github.com/kardianos/govendor \
    && govendor sync -v  \ 
    &&  go build -o cosme . && ls -la cosme\
    && echo "completed !!!"


FROM golang:1.11-alpine
WORKDIR /go/src/
COPY --from=buildgolang  /go/src/cosme/cosme /go/src/cosme 
CMD [ "./cosme" ]