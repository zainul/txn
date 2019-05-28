FROM golang:1.10.3

RUN apt-get update
RUN apt-get install sudo

# env system
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# env app

# set local time
RUN ln -sf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
RUN echo "Asia/Jakarta" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir -p /go/src/github.com/zainul/txn

# add current porject
ADD . /go/src/github.com/zainul/txn

# set working dir
WORKDIR /go/src/github.com/zainul/txn

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v --vendor-only

# install the migrations
RUN go get -u github.com/zainul/gan
RUN go get github.com/jinzhu/gorm
RUN go get github.com/c-bata/go-prompt
RUN cd /go/src/github.com/zainul/gan/ && go build 
RUN cp /go/src/github.com/zainul/gan/gan /usr/bin/

# build the binary
RUN go build /go/src/github.com/zainul/txn/cmd/main.go

CMD ["/go/src/github.com/zainul/txn/cmd/main"]
