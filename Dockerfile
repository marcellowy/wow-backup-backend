FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

CMD ["app","-c","/usr/local/app-config/config.yaml"]

# docker build -t my-golang-app .
# docker run -it --rm -p 44889:8080 -v /data/etc/wow-backup:/usr/local/app-config --name my-running-app my-golang-app