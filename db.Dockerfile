FROM docker.io/library/postgres:14.4-bullseye
EXPOSE 5432
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_USER=goapp
ADD ./db_init.sql /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*
