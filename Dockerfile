FROM golang

RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/dongjujang/slackchat-cus

ADD . /go/src/test/crawler_slack
RUN go install test/crawler_slack

EXPOSE 8888

ENTRYPOINT ["/go/bin/crawler_slack"]