FROM migrate/migrate:4
WORKDIR /app
COPY migrations /db/migrations

CMD ["/bin/sh"]