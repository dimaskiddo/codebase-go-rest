FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

ARG SERVICE_NAME="codebase-go-rest"
ENV CONFIG_ENV="PROD"

WORKDIR /usr/src/app

COPY misc/ ./misc
COPY dist/${SERVICE_NAME}_linux_amd64/main ./main

RUN chmod 777 misc/stores misc/uploads

EXPOSE 3000
HEALTHCHECK --interval=5s --timeout=3s CMD ["curl", "http://127.0.0.1:3000/health"] || exit 1

VOLUME ["/usr/src/app/misc/stores","/usr/src/app/misc/uploads"]
CMD ["./main"]
