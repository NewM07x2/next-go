 => ERROR [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  0.3s
------
 > [go-app builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api:
0.238 internal/router/router.go:7:2: package next-go-sample/pkg/graph/generated is not in std (/usr/local/go/src/next-go-sample/pkg/graph/generated)
------
failed to solve: process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api" did not complete successfully: exit code: 1