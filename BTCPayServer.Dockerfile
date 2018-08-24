FROM golang:alpine as builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Install dependencies and build the binaries.
RUN apk add --no-cache \
    git \
    make

WORKDIR /go/src/github.com/BTCGPU/lnd
COPY . .

RUN make \
&&  make install

# Start a new, final image.
FROM alpine as final

# Add bash and ca-certs, for quality of life and SSL-related reasons.
RUN apk --no-cache add \
    bash \
    ca-certificates

ENV LND_DATA /data
ENV LND_BITCOIND /deps/.bitcoingold
ENV LND_BTCD /deps/.btcd

RUN mkdir "$LND_DATA" && mkdir "/deps" && mkdir "$LND_BITCOIND" && mkdir "$LND_BTCD" && \
    ln -sfn "$LND_DATA" /root/.lnd && ln -sfn "$LND_BITCOIND" /root/.bitcoingold && ln -sfn "$LND_BTCD" /root/.btcd

# Define a root volume for data persistence.
VOLUME /data

# Copy the binaries from the builder image.
COPY --from=builder /go/bin/lncli /bin/
COPY --from=builder /go/bin/lnd /bin/

COPY docker-entrypoint.sh /docker-entrypoint.sh
# Specify the start command and entrypoint as the lnd daemon.
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["lnd"]