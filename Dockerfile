FROM jrudio/go-plex-client-base:test

WORKDIR /app/go-plex-client

# copy source code to docker container
COPY ./ /app/go-plex-client

# copy start up script
COPY ./cloud-build/pms-startup.sh /app/go-plex-client

RUN go mod tidy

ENTRYPOINT [ "/bin/sh", "./pms-startup.sh" ]