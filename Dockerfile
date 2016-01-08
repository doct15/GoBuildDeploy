
# DOCKER-VERSION 0.3.4

FROM ubuntu:12.04

RUN apt-get -y update

ADD ./bin/server /

EXPOSE 8080
CMD ["/server", "8080"]
