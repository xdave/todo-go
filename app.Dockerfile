FROM docker.io/library/golang:1.17.0-buster
WORKDIR /home
COPY . .
RUN go mod download
RUN go build -o ../output main.go
EXPOSE 8000
ENTRYPOINT ["../output"]