FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

CMD ["app","-c","/usr/local/app-config/config.yaml"]

# docker build --rm -t wow-backup .
# docker run -itd -p 31445:8080 \
# -v /data/etc/wow-backup:/usr/local/app-config \
# -v /var/log/wow-backup:/var/log \
# -v /data/wow-backup:/data
# --name wow-backup wow-backup