#!/bin/bash
# user_test
go test -v tests/models/user/user_test.go \
	-dbName="auth" -network="127.0.0.1:3306" \
	-dbUser="docker" -dbPassword="docker"
#user benchmark
go test -bench=. tests/models/user/bench_test.go \
	-dbName="auth" -network="127.0.0.1:3306" \
	-dbUser="docker" -dbPassword="docker"

#auth_test
go test -v tests/auth/auth_test.go \
	-dbName="auth" -network="127.0.0.1:3306" \
	-dbUser="root" -dbPassword="root"
# client test
go test -v tests/client/client_test.go \
	-dbName="auth" -network="127.0.0.1:3306" \
	-dbUser="root" -dbPassword="root"