FROM jrudio/go-plex-client-base:test

WORKDIR /app/go-plex-client

ARG MEDIA_PATH=/data

# copy source code to docker container
COPY ./ /app/go-plex-client

COPY ./cloud-build/media/love_shouldnt_hurt.mp4 .
COPY ./cloud-build/media/media_generator.sh .

# create fake movie and tv show directories
RUN ./media_generator.sh ./love_shouldnt_hurt.mp4 ${MEDIA_PATH}

# copy start up script
COPY ./cloud-build/pms-startup.sh /app/go-plex-client

RUN go mod tidy

ENTRYPOINT [ "/bin/sh", "./pms-startup.sh" ]