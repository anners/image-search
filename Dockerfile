FROM golang:1.10

MAINTAINER Ann Wallace annerz@gmail.com

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/anners/image-service

# Build the cat-service command inside the container.
RUN go install github.com/anners/image-service

# Run the cat-service command by default when the container starts.
ENTRYPOINT /go/bin/image-service

# Document that the service listens on port 8080.
EXPOSE 8888
