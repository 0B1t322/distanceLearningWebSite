module github.com/0B1t322/auth-service

go 1.15

replace (
	github.com/0B1t322/distanceLearningWebSite/pkg/db => ../pkg/db
)

require (
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/google/grpc v1.34.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20201231184435-2d18734c6014 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
	github.com/0B1t322/distanceLearningWebSite/pkg/db v1.0.0
)

