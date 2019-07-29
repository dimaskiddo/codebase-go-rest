FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

ARG SERVICE_NAME="codebase-go-rest"
ENV CONFIG_ENV="production"

WORKDIR /usr/src/app

COPY share/ ./share
COPY dist/${SERVICE_NAME}_linux_amd64/main ./main

RUN chmod 777 share/stores share/uploads

EXPOSE 3000
HEALTHCHECK --interval=5s --timeout=3s CMD ["curl", "http://127.0.0.1:3000/health"] || exit 1

VOLUME ["/usr/src/app/share/stores","/usr/src/app/share/uploads"]
CMD ["./main"]
