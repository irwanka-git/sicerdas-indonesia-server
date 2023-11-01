FROM alpine:3.14
RUN apk add libc6-compat
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

COPY /cron/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh
CMD ["/entrypoint.sh"]