# Here we are getting our Go as Base image
FROM golang:1.18-alpine
RUN apk update && apk add --no-cache gcc musl-dev git bash
# create user in the docker image inside Alpine
RUN addgroup -S gogroup
RUN adduser -G gogroup -S gouser
USER gouser
# Creating a new directory for app files and setting path in the container
RUN mkdir -p /home/gouser/app && chown -R gouser:gogroup /home/gouser/app
# setting working directory in the container
WORKDIR /home/gouser/app
# grant permission of the project directory to gouser user
COPY --chown=gouser:gouser . .
# installing the dependencies into the container
# pre-copy/cache go.mod for pre-downloading dependencies
# and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
# build the app
COPY . .
RUN go build -ldflags '-w -s' -a -o /home/gouser/server ./cmd/app
# container exposed network port number
EXPOSE $PORT
# command to run within the container
CMD ["/home/gouser/server"]
