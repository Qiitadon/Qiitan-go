ARG BASEIMAGE="golang:alpine"

# -----------------------------------------------------------------------------
#  Homebrew Stage
# -----------------------------------------------------------------------------
FROM homebrew/brew:latest AS brew

WORKDIR /home/linuxbrew

RUN \
    # Install qiitan
    brew install qithub-bot/apps/qiitan && \
    # Smoke-test
    qiitan -version

RUN \
    # Download Tengo
    curl -sSL \
    https://github.com/d5/tengo/releases/download/v2.10.0/tengo_2.10.0_linux_amd64.tar.gz \
    -o tengo.tar.gz && \
    tar -xvf ./tengo.tar.gz && \
    # Smoke-test
    ./tengo -version

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM ${BASEIMAGE}

# Install qiitan
COPY --from=brew /home/linuxbrew/.linuxbrew/bin/qiitan /usr/local/bin/qiitan
# Install tengo
COPY --from=brew /home/linuxbrew/tengo /usr/local/bin/tengo

# Smoke-tests
RUN \
    qiitan -version && \
    tengo -version

# Install perf
RUN \
    apk add --no-cache perf && \
    # Smoke-test
    perf --version
