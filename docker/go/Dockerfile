## ------------
##	Development
## ------------

FROM golang:1.21-alpine as development

WORKDIR /var/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/psampaz/go-mod-outdated@latest

COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]

## --------
##	Builder
## --------

FROM golang:1.21-alpine as builder

WORKDIR /var/app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./bin/ciby ./src/go

## -----------
##	Production
## -----------

FROM gcr.io/distroless/static-debian11 as production

COPY --from=builder /var/app/bin ./bin
COPY --from=builder /var/app/config ./config
COPY --from=builder /var/app/public ./public
COPY --from=builder /var/app/ui/jinja ./ui/jinja

CMD ["./bin/ciby"]
