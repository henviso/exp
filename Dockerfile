FROM alpine
COPY release/kongdesk /bin/
EXPOSE 8080
ENTRYPOINT ["kongdesk"]
