# Builder Image
# ---------------------------------------------------
FROM dimaskiddo/alpine:go-1.12 AS go-builder

WORKDIR /usr/src/app

COPY . ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -o main cmd/main/main.go


# Final Image
# ---------------------------------------------------
FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

ARG SERVICE_NAME="codebase-go-rest"
ENV PATH="$PATH:/usr/app/${SERVICE_NAME}" \
    CONFIG_ENV="production"

WORKDIR /usr/app/${SERVICE_NAME}

COPY --from=go-builder /usr/src/app/config/ ./config
COPY --from=go-builder /usr/src/app/main ./main

RUN chmod 777 config/stores config/uploads

EXPOSE 3000
HEALTHCHECK --interval=5s --timeout=3s CMD ["curl", "http://127.0.0.1:3000/health"] || exit 1

VOLUME ["/usr/app/config/stores","/usr/app/config/uploads"]
CMD ["main"]
