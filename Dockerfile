FROM golang:1.11-alpine

ENV GOPATH=/go
WORKDIR ${GOPATH}/src/cosme

COPY . .

RUN sh bin/install.sh

CMD ["./cosme"]

EXPOSE 3000