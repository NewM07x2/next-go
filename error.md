 => ERROR [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  0.3s
------
 > [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api:
0.238 internal/router/router.go:7:2: package next-go-sample/pkg/graph/generated is not in std (/usr/local/go/src/next-go-sample/pkg/graph/generated)
------
failed to solve: process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api" did not complete successfully: exit code: 1




go run github.com/99designs/gqlgen generate

(base) mnitta@MasatonoMacBook-Pro go-app % go run github.com/99designs/gqlgen generate
../../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/code/packages.go:12:2: missing go.sum entry for module providing package golang.org/x/tools/go/packages (imported by github.com/99designs/gqlgen/codegen/config); to add:
        go get github.com/99designs/gqlgen/codegen/config@v0.17.81
../../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/imports/prune.go:13:2: missing go.sum entry for module providing package golang.org/x/tools/go/ast/astutil (imported by github.com/99designs/gqlgen/internal/imports); to add:
        go get github.com/99designs/gqlgen/internal/imports@v0.17.81
../../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/imports/prune.go:14:2: missing go.sum entry for module providing package golang.org/x/tools/imports (imported by github.com/99designs/gqlgen/api); to add:
        go get github.com/99designs/gqlgen/api@v0.17.81
../../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/main.go:17:2: missing go.sum entry for module providing package github.com/urfave/cli/v2 (imported by github.com/99designs/gqlgen); to add:
        go get github.com/99designs/gqlgen@v0.17.81


go get github.com/99designs/gqlgen/codegen/config@v0.17.81
go get github.com/99designs/gqlgen/internal/imports@v0.17.81
go get github.com/99designs/gqlgen/api@v0.17.81
go get github.com/99designs/gqlgen@v0.17.81

go mod tidy
go run github.com/99designs/gqlgen generate

(base) mnitta@MasatonoMacBook-Pro go-app % go run github.com/99designs/gqlgen generate

(base) mnitta@MasatonoMacBook-Pro go-app % go get github.com/99designs/gqlgen/codegen/config@v0.17.81
go get github.com/99designs/gqlgen/internal/imports@v0.17.81
go get github.com/99designs/gqlgen/api@v0.17.81
go get github.com/99designs/gqlgen@v0.17.81
(base) mnitta@MasatonoMacBook-Pro go-app % go run github.com/99designs/gqlgen generate
unable to open schema: open schema.graphql: no such file or directory
exit status 1


(base) mnitta@MasatonoMacBook-Pro go-app % go run github.com/99designs/gqlgen generate
unable to open schema: open schema.graphql: no such file or directory
exit status 1

cd /Users/mnitta/devlop/next-go/go-app
go run github.com/99designs/gqlgen generate --config pkg/graph/gqlgen.yml


 => ERROR [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  0.4s
 => CACHED [next-app 2/4] WORKDIR /app                                                      0.0s
 => CACHED [next-app 3/4] COPY package*.json ./                                             0.0s
 => CACHED [next-app 4/4] COPY . .                                                          0.0s
 => [next-app] exporting to image                                                           0.0s
 => => exporting layers                                                                     0.0s
 => => writing image sha256:078b8979098aad7635e21ca908a0137f9a845e4235c22ae607b48476a7d977  0.0s
 => => naming to docker.io/library/next-go-next-app                                         0.0s
 => [next-app] resolving provenance for metadata file                                       0.0s
------
 > [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api:
0.257 internal/router/router.go:7:2: package next-go-sample/pkg/graph/generated is not in std (/usr/local/go/src/next-go-sample/pkg/graph/generated)
------
failed to solve: process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api" did not complete successfully: exit code: 1