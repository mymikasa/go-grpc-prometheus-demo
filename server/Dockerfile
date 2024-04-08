FROM golang:1.22.1

USER root
RUN mkdir /src

COPY ./server /src/server

EXPOSE 8089

ENTRYPOINT [ "/src/server" ]