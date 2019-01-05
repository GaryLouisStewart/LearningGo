# enable dynamic setting of docker image and version (overridable on buildtime)
# docker build --build-arg BUILDTIME_IMAGE=[CHANGEME] --build-arg BUILDTIME_VERSION=[CHANGEME] --name $DKR_IMG_NAME .

ARG $BUILDTIME_IMAGE=golang
ARG $BUILDTIME_VERSION=latest
ENV IMAGE=$BUILDTIME_IMAGE
ENV VERSION=$BUILDTIME_VERSION

FROM $IMAGE:$VERSION

RUN mkdir -p /opt/mock-api

COPY ./mock-api /opt/mock-api

COPY ./docker-entrypoint.sh /docker-entrypoint.sh

CMD ["start"]

ENTRYPOINT ["/docker-entrypoint.sh"]
