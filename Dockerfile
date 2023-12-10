# Dockerfile definition for Backend application service.

# From which image we want to build. This is basically our environment.
FROM golang:1.19-alpine as Build

# Set the working directory to the root of your application
WORKDIR /go/src/github.com/bagusandrian/sawitpro

# Copy all files to the working directory
COPY . .

# Build our binary at root location.
RUN go build -o /main cmd/main.go

####################################################################
# This is the actual image that we will be using in production.
FROM alpine:latest

# Set the working directory
WORKDIR /app

# We need to copy the binary from the build image to the production image.
COPY --from=Build /main .

# Copy the config file to the expected location
COPY files/etc/config.yaml /etc/sawitpro/

# Debugging statements
RUN ls -l
RUN ls -l /etc/sawitpro/
RUN cat /etc/sawitpro/config.yaml
ENV APPS_ENV=DOCKER
RUN echo "env $APPS_ENV"

# This is the port that our application will be listening on.
EXPOSE 8002

# This is the command that will be executed when the container is started.
CMD ["./main"]
