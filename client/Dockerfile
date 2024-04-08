FROM golang:1.22.1

USER root
RUN mkdir /src

COPY ./client /src/client

EXPOSE 8081

ENTRYPOINT [ "/src/client" ]