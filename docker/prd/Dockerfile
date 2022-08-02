# start from golang
FROM golang:1-buster as builder
WORKDIR /app

# copy dependencies info
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source
COPY . .

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dota2 /app/cmd/dota2/main.go

# now move to scratch
FROM scratch
WORKDIR /app

# copy the binary from builder
COPY --from=builder /app/dota2 .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# run the binary
ENTRYPOINT ["./dota2"]
