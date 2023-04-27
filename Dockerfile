FROM golang:1.20-alpine as builder

ARG TARGETPLATFORM
ARG TARGETARCH
RUN echo building for "$TARGETPLATFORM"

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH GO111MODULE=on go build -a -o telegram-gpt-bot ./cmd/server

FROM alpine

COPY --from=builder /workspace/telegram-gpt-bot /telegram-gpt-bot

ENV TELEGRAM_BOT_TOKEN "TELEGRAM_BOT_TOKEN"
ENV OPENAI_API_KEY "OPENAI_API_KEY"

# Run the binary.
ENTRYPOINT ["/telegram-gpt-bot"]