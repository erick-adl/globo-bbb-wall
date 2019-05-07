# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Erick Lima <erick.adlima@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/erick-adl/globo-bbb-wall/backend_go

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN mkdir -p $HOME/.aws

#add access sqs
RUN echo -e "[default]\naws_access_key_id = ********************\naws_secret_access_key = ****************************************" >> $HOME/.aws/credentials

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the executable
CMD ["backend_go"]