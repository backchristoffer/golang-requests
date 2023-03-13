FROM golang:1.16-alpine
WORKDIR /app
COPY . ./
COPY go.mod .
RUN go mod download
COPY *.go .
RUN go version
RUN go build -o golang-requests
CMD ["/app/golang-requests"]
