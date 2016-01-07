
# DOCKER-VERSION 0.3.4

FROM ubuntu:12.04

RUN sudo apt-get -y update

ADD . /

EXPOSE 8080
CMD ["/bin/server", "8080"]
 
