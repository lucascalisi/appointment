FROM golang:latest

EXPOSE 8080
ENV HOST=0.0.0.0
ENV PORT=8080

ENV DBPASSWORD=
ENV DBNAME=
ENV DBHOST=localhost
ENV DBPORT=3306
ENV DBUSER=root

WORKDIR /src
COPY . .
#ADD http/cmd/a-p-pointment-backend-server/main /app/main
#CMD ["/app/main",  "--host", "0.0.0.0", "--port", "8080"]
