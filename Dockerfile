FROM golang:1.16 AS builder

WORKDIR /feather

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build
#RUN make tests

# Runner container

FROM alpine:3.13 AS runner
WORKDIR /feather
COPY --from=builder /feather/service-core /feather/service-core
EXPOSE 8080

ENV PORT=8080
ENV HOST=0.0.0.0

ENTRYPOINT ["./service-core"]