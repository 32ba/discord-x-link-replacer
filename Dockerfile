FROM golang:1.21.4 AS builder
WORKDIR /app
ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o discord-x-link-replacer
RUN strip discord-x-link-replacer

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=builder /app/discord-x-link-replacer /discord-x-link-replacer
USER nonroot
CMD ["/discord-x-link-replacer"]