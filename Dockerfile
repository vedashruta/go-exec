FROM golang:1.22.4
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server
EXPOSE 9099
CMD ["./server"]