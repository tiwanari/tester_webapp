# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

# Set go bin which doesn't appear to be set already
ENV GOBIN /go/bin

WORKDIR /app

# Set an env var that matches your github repo name
ENV SRC_DIR=/go/src/github.com/tiwanari/tester_webapp/

# Add the source code
ADD . $SRC_DIR

# Build it
RUN go get -u github.com/golang/dep/cmd/dep/...; \
    cd $SRC_DIR/tester; \
    make; \
    cp webapp /app/; \
    ln -s $SRC_DIR/tester/src/webapp/views /app/views

# For `docker run`
ENTRYPOINT ["./webapp"]
