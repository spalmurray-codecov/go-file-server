FROM scratch

COPY --chmod=400 config/passwd /etc/passwd
COPY --chmod=444 www /var/www
COPY --chmod=111 go-file-server /

EXPOSE 80

USER www-data

CMD ["/go-file-server"]
