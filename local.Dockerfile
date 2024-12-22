FROM postgres:15

# Set environment variables (these can be overridden in docker-compose.yml)
ENV POSTGRES_DB=mydb
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword

# Copy any initialization scripts (optional)
COPY ./init.sql /docker-entrypoint-initdb.d/