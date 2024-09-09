# GENERATE GO BINARY
FROM golang:alpine as builder

# Menyalin kode dari host dan melakukan kompilasi
WORKDIR $GOPATH/src/bitbucket.org/ms-sv-jira
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /ms-sv-jira

# RUNNING GO BINARY
FROM alpine:3.8

# Menambahkan curl dan openssl
RUN apk --no-cache add curl openssl

# Mengatur timezone ke Asia/Jakarta
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

# Generate kunci privat dan kunci publik SSL
RUN mkdir -p /ssl \
    && openssl genrsa -out /ssl/privkey.pem 4096 \
    && openssl rsa -in /ssl/privkey.pem -out /ssl/pubkey.pem -pubout -outform PEM

# Menyalin binary Go dari tahap kompilasi
COPY --from=builder /ms-sv-jira ./

# Opsi untuk menyalin file tambahan
COPY . .

# Menjalankan aplikasi
ENTRYPOINT ["/ms-sv-jira"]