FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

COPY build/ /opt/app/
WORKDIR /opt/app

EXPOSE 3000
CMD ["/opt/app/main"]
