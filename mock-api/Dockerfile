FROM golang:1.13-alpine

RUN mkdir -p /opt/mock-api

COPY ./mock-api /opt/mock-api

COPY ./docker-entrypoint.sh /docker-entrypoint.sh

EXPOSE 1111

CMD ["start"]

ENTRYPOINT ["/docker-entrypoint.sh"]
