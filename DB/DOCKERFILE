# FROM postgres:latest
FROM --platform=linux/amd64 postgres:latest

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]