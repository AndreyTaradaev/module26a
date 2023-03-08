FROM golang:1.20
# RUN mkdir -p /go/src/module26a
WORKDIR /go/src/module26a
RUN go env -w GO111MODULE=auto
RUN git clone https://github.com/AndreyTaradaev/module26a.git /go/src/module26a
RUN go install .

FROM alpine:latest
LABEL version="v1.0(thin)"
LABEL maintainer="AndreyTaradaev<test@test.ru>"
WORKDIR /root/
COPY --from=0 /go/bin/pipeline .
ENTRYPOINT ./pipeline -debug





