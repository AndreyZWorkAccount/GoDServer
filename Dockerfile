FROM alpine:3.3

COPY ./graph /opt/bin/

CMD [ "/opt/bin/" ]