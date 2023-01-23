FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/edwynpratama1/serverLocalGo
RUN cd /build && git clone https://github.com/edwynpratama1/serverLocalGo.git

RUN cd /build/serverLocalGo && go build

EXPOSE 8080

ENTRYPOINT [ "/build/serverLocalGo" ]