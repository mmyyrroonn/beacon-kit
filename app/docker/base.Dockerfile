# syntax=docker/dockerfile:1
#
# Copyright (C) 2022, Berachain Foundation. All rights reserved.
# See the file LICENSE for licensing terms.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#######################################################
###           Stage 0 - Build Arguments             ###
#######################################################

ARG GO_VERSION=1.21.5
ARG RUNNER_IMAGE=alpine
ARG BUILD_TAGS="netgo,ledger,muslc"
ARG GOARCH=amd64
ARG GOOS=linux
ARG NAME=beacond
ARG APP_NAME=beacond
ARG DB_BACKEND=pebbledb
ARG CMD_PATH=./app/beacond
ARG FOUNDRY_DIR=contracts

#######################################################
###         Stage 1 - Build the Application         ###
#######################################################

FROM golang:${GO_VERSION}-alpine3.18 as builder

ARG GIT_VERSION
ARG GIT_COMMIT
ARG BUILD_TAGS

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

# Setup some alpine stuff that nobody really knows how or why it works.
RUN set -eux; \
    apk add --no-cache git linux-headers ca-certificates build-base
# Set the working directory
WORKDIR /workdir

# Copy the go.mod and go.sum files for each module
COPY ./go.mod ./go.sum ./

# Download the go module dependencies
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy the rest of the source code
COPY . .

# Build args
ARG NAME
ARG GOARCH
ARG GOOS
ARG APP_NAME
ARG DB_BACKEND
ARG CMD_PATH

# Build beacond
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    env GOOS=${GOOS} GOARCH=${GOARCH} && \
    env NAME=${NAME} DB_BACKEND=${DB_BACKEND} && \
    env APP_NAME=${APP_NAME} && \
    go build \
    -mod=readonly \
    -tags ${BUILD_TAGS} \
    -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=${NAME} \
    -X github.com/cosmos/cosmos-sdk/version.AppName=${APP_NAME} \
    -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -X github.com/cosmos/cosmos-sdk/types.DBBackend=$DB_BACKEND \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /workdir/build/bin/ \
    ${CMD_PATH}

#######################################################
###        Stage 2 - Prepare the Final Image        ###
#######################################################

FROM ${RUNNER_IMAGE}

# Build args
ARG APP_NAME

# Copy over built executable into a fresh container.
COPY --from=builder /workdir/build/bin/${APP_NAME} /bin/