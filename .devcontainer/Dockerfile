# =============================================================================
#  Development Container for Go + VSCode + Remote-Containers.
#  (Works for GitHub Codespaces as well)
# =============================================================================
# [Choice] https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go
ARG VARIANT="latest"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM ghcr.io/keinos/vscode-dev-container-go:${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# USER root
# RUN apk add --no-cache <your-package-list-here>

# [Optional] Uncomment this section to go install anything else you need.
USER vscode
RUN cd /tmp && \
    git config --global init.defaultBranch main && \
    go install "github.com/d5/tengo/cmd/tengo@latest" && \
    go install "golang.org/x/lint/golint@latest" && \
    go install "github.com/ekalinin/github-markdown-toc.go/cmd/gh-md-toc@latest" && \
    mkdir -p ~/.vscode-server/extensions
