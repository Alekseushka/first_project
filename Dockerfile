# Use the official PostgreSQL image as the base
FROM postgres:latest

# Set environment variables for PostgreSQL (optional, can be passed during docker run)
ENV POSTGRES_DB=mydatabase
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mysecretpassword