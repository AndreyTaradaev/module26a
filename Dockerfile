FROM golang:1.20
LABEL version="1.0.0"
LABEL maintainer="AndreyTaradaev<test@test.ru>"
# RUN mkdir -p /go/src/module26a
WORKDIR /go/src/module26a
RUN go env -w GO111MODULE=auto
RUN git clone https://github.com/AndreyTaradaev/module26a.git /go/src/module26a
RUN  ls   ./
RUN go install .
ENTRYPOINT /go/bin/pipeline -debug
