#FROM golang:latest
#
#RUN apt-get update && apt-get install -y \
#    git \
#    curl \
#    build-essential
#
#WORKDIR /app
#
#COPY . /app
#
#RUN go get -d -v ./...
#
#RUN go build -o main .
#
#EXPOSE 8080
#
#CMD ["./main"]
FROM node:18-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["./main"]
EXPOSE 80