# Build the application
FROM --platform=$BUILDPLATFORM golang:1.21.3-bookworm as base
ARG build_tags=""

WORKDIR /build

COPY . .
RUN go mod download

ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS
RUN echo "TARGETARCH: $TARGETARCH"; \
    echo "TARGETOS: $TARGETOS"; \
    echo "TARGETPLATFORM: $TARGETPLATFORM";
RUN GOARCH=$TARGETARCH GOOS=$TARGETOS go build -tags="${build_tags}" -o ./${{ application_name }} main.go

# Run the application
FROM alpine:3.18.4

RUN apk add --no-cache libstdc++ gcompat libgomp; \
    apk add --update busybox>1.3.1-r0; \
    apk add --update openssl>3.1.4-r1

RUN apk add doas; \
    adduser -S dommyuser -D -G wheel; \
    echo 'permit nopass :wheel as root' >> /etc/doas.d/doas.conf;
RUN chmod g+rx,o+rx /

COPY --from=base /build/${{ application_name }} /app/${{ application_name }}

ENV LOG_ENV=production

# Command to run
ENTRYPOINT ["/app/${{ application_name }}"]
