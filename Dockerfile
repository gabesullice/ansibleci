FROM alpine

WORKDIR /app

RUN apk update \
    && apk --update add \
    go \
    make \
    openssh \
    python \
    python-dev \
    build-base \
    py-pip

RUN pip install ansible

ADD https://github.com/elevatedthird/deploykit/archive/master.tar.gz /app/master.tar.gz
RUN tar -xzf /app/master.tar.gz -C /app/

ADD src /app/src

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY settings /app/settings

COPY src /go/src/ansibleci

RUN go install ansibleci
