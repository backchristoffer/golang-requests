FROM golang:latest
WORKDIR /app
COPY . ./
RUN go mod download
RUN go version
RUN go build
CMD ["/golang-requests"]