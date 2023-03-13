FROM golang:latest
WORKDIR /app
COPY . ./
ENV GOOS=linux
ENV GOARCH=arm
RUN go mod download
RUN go version
RUN go build
CMD ["/golang-requests"]
