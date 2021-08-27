# Copyright (c) 2021 KEINOS and The Sqiima Contributors (https://github.com/KEINOS/sqiima/graphs/contributors)
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# =============================================================================
#  Alpine-based Dockerfile of Go v1.16 for the devcontainer
# =============================================================================
#  This container is intended to be used for development and testing which
#  includes most of the latest packages and tools needed to get started.

# [Choice] https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go
ARG VARIANT="latest"
FROM ghcr.io/keinos/vscode-dev-container-go:${VARIANT}

USER root

# [Optional] Uncomment this section to install additional OS packages.
# RUN apk add --no-cache <your-package-list-here>

USER vscode

# [Optional] Uncomment the next line to go install anything else you need
# RUN cd /tmp && go intall <your-dependency-or-tool>@<version>
