FROM golang:1.20-alpine

WORKDIR /var/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]
