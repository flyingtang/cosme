FROM golang:1.11-alpine as buildgolang


WORKDIR /go/src/cosme

COPY . .

RUN apk add git \
    && rootdir=`pwd` \
    && golangpath="$rootdir/vendor/golang.org/x"  \
    && rm -rf $golangpath \
    && mkdir -p $golangpath && ls -la $golangpath && cd $golangpath \
    && git clone https://github.com/golang/crypto.git && git clone https://github.com/golang/sys.git \
    && cd $rootdir \
    && go get -u github.com/kardianos/govendor \
    && govendor sync -v  \ 
    &&  go build -o app . && ls -la cosme\
    && echo "completed !!!"
# 清除工作

# FROM golang:1.11-alpine
# ENV GOPATH /go

# WORKDIR $GOPATH/src/cosme
# COPY --from=buildgolang  /go/src/cosme/app /go/src/cosme

CMD [ "./cosme" ]