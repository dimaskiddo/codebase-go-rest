FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

WORKDIR /opt/app
COPY build/ .
RUN chmod 777 uploads

EXPOSE 3000
CMD ["/opt/app/main"]
