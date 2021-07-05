FROM golang:1.16-alpine as build
RUN apk --no-cache add tzdata
WORKDIR /app
ADD go-backups .
ADD config ./config

FROM alpine as final
WORKDIR /app
COPY --from=build /app/go-backups .
COPY --from=build /app/config ./config
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
ENV FLAG=""
ENTRYPOINT /app/go-backups ${FLAG}