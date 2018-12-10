FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

WORKDIR /opt/app
COPY build/ .
RUN chmod 777 uploads

EXPOSE 3000
HEALTHCHECK --interval=3s --timeout=3s CMD ["curl", "http://127.0.0.1:3000/health"] || exit 1
CMD ["/opt/app/main"]
