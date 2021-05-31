FROM postgres
WORKDIR /docker-entrypoint-initdb.d
ADD ./sql/script.sql /docker-entrypoint-initdb.d
EXPOSE 5432