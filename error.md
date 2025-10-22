(base) mnitta@MasatonoMacBook-Pro next-go % docker-compose up --build -d           
WARN[0000] /Users/mnitta/devlop/next-go/docker-compose.yml: the attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion 
Compose can now delegate builds to bake for better performance.
 To do so, set COMPOSE_BAKE=true.
[+] Building 2.6s (22/25)                                                   docker:desktop-linux
 => [next-app internal] load build definition from Dockerfile.frontend                      0.0s
 => => transferring dockerfile: 226B                                                        0.0s
 => [go-app internal] load build definition from Dockerfile.backend                         0.0s
 => => transferring dockerfile: 526B                                                        0.0s
 => [next-app internal] load metadata for docker.io/library/node:20                         1.1s
 => [go-app internal] load metadata for docker.io/library/alpine:latest                     1.1s
 => [go-app internal] load metadata for docker.io/library/golang:1.23-alpine                1.1s
 => [next-app internal] load .dockerignore                                                  0.0s
 => => transferring context: 2B                                                             0.0s
 => [go-app internal] load .dockerignore                                                    0.0s
 => => transferring context: 335B                                                           0.0s
 => [next-app 1/4] FROM docker.io/library/node:20@sha256:8fe515fa9035879b707797479df19e741  0.0s
 => => resolve docker.io/library/node:20@sha256:8fe515fa9035879b707797479df19e7410947a6675  0.0s
 => [next-app internal] load build context                                                  0.0s
 => => transferring context: 11.39kB                                                        0.0s
 => [go-app builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:383395b794dffa5b  0.0s
 => [go-app stage-1 1/4] FROM docker.io/library/alpine:latest@sha256:4b7ce07002c69e8f3d704  0.0s
 => [go-app internal] load build context                                                    0.0s
 => => transferring context: 3.56kB                                                         0.0s
 => CACHED [go-app stage-1 2/4] RUN apk --no-cache add ca-certificates                      0.0s
 => CACHED [go-app stage-1 3/4] WORKDIR /root/                                              0.0s
 => CACHED [next-app 2/4] WORKDIR /app                                                      0.0s
 => CACHED [next-app 3/4] COPY package*.json ./                                             0.0s
 => CACHED [next-app 4/4] COPY . .                                                          0.0s
 => [next-app] exporting to image                                                           0.0s
 => => exporting layers                                                                     0.0s
 => => writing image sha256:014e9ae161e107fa0915d33b62ea4e872de4b663f18590063fe0450b5b381b  0.0s
 => => naming to docker.io/library/next-go-next-app                                         0.0s
 => CACHED [go-app builder 2/6] WORKDIR /app                                                0.0s
 => [go-app builder 3/6] COPY go.mod go.sum ./                                              0.0s
 => ERROR [go-app builder 4/6] RUN go mod download                                          1.4s
 => [next-app] resolving provenance for metadata file                                       0.0s
------
 > [go-app builder 4/6] RUN go mod download:
1.359 go: github.com/99designs/gqlgen@v0.17.81: module github.com/99designs/gqlgen@v0.17.81 requires go >= 1.24.0 (running go 1.23.12; GOTOOLCHAIN=local)
------
failed to solve: process "/bin/sh -c go mod download" did not complete successfully: exit code: 1
(base) mnitta@MasatonoMacBook-Pro next-go %
