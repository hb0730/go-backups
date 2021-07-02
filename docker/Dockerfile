FROM alpine
WORKDIR /app
RUN mkdir ./config
ENV FLAG=""
COPY ./go-backups .
COPY ./config/*.yaml ./config
RUN  ls -al
RUN cd config
RUN ls -al
ENTRYPOINT /app/go-backups ${FLAG}