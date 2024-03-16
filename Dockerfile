FROM golang:alpine AS build

# Get args and environment vars
ARG PORT=3002
ENV PORT=$PORT

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o /tmp/app

# This results in a single layer image
FROM alpine:edge

# Get args and environment vars
ARG PORT=3002
ENV PORT=$PORT

COPY --from=build /tmp/app /sbin/app

# Expose the port
EXPOSE $PORT

CMD /sbin/app -port $PORT