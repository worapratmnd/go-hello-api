FROM golang:1.20 as builder

# Set destination for COPY
WORKDIR /app

ARG APPNAME=hello
ENV ENV_APPNAME=$APPNAME

# Download Go modules
# COPY . /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY . ./

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /$APPNAME

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 3001

# second state
FROM scratch

COPY --from=builder /hello .

EXPOSE 3001

# Run
CMD ["./hello"]