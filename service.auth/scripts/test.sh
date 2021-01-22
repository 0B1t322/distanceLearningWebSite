#!/bin/bash

dbName="auth"
dbUser="root"
dbPassword="root"
network="127.0.0.1:3306"

# user_test
# go test -v tests/models/user/user_test.go \
# 	-dbName=$dbName -network=$network \
# 	-dbUser=$dbUser -dbPassword=$dbUser
#user benchmark
go test -bench=. tests/models/user/bench_test.go 

#auth_test
go test -v tests/auth/auth_test.go 
# client test
go test -v tests/client/client_test.go 