 => [go-app stage-1 3/4] WORKDIR /root/                                                     0.0s
 => [go-app builder 5/7] COPY . .                                                           0.0s
 => ERROR [go-app builder 6/7] RUN go run github.com/99designs/gqlgen generate --config pk  0.3s
------
 > [go-app builder 6/7] RUN go run github.com/99designs/gqlgen generate --config pkg/graph/gqlgen.yml:
0.256 /go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/code/packages.go:12:2: missing go.sum entry for module providing package golang.org/x/tools/go/packages (imported by github.com/99designs/gqlgen/codegen/config); to add:
0.256   go get github.com/99designs/gqlgen/codegen/config@v0.17.81
0.256 /go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/imports/prune.go:13:2: missing go.sum entry for module providing package golang.org/x/tools/go/ast/astutil (imported by github.com/99designs/gqlgen/internal/imports); to add:
0.256   go get github.com/99designs/gqlgen/internal/imports@v0.17.81
0.256 /go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/internal/imports/prune.go:14:2: missing go.sum entry for module providing package golang.org/x/tools/imports (imported by github.com/99designs/gqlgen/api); to add:
0.256   go get github.com/99designs/gqlgen/api@v0.17.81
0.256 /go/pkg/mod/github.com/99designs/gqlgen@v0.17.81/main.go:17:2: missing go.sum entry for module providing package github.com/urfave/cli/v2 (imported by github.com/99designs/gqlgen); to add:
0.256   go get github.com/99designs/gqlgen@v0.17.81
------
failed to solve: process "/bin/sh -c go run github.com/99designs/gqlgen generate --config pkg/graph/gqlgen.yml" did not complete successfully: exit code: 1