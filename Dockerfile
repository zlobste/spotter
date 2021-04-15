FROM golang:1.16

# Define workdir
RUN mkdir ./app
COPY . ./app
WORKDIR ./app

# Download dependencies
RUN go get -d -v

# Install the package
RUN go install -v

# Set config file
ENV CONFIG=config.yaml

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["spotter"]