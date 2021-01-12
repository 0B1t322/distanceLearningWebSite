module github.com/0B1t322/distanceLearningWebSite

go 1.15

// replace (
// 	github.com/0B1t322/distanceLearningWebSite/pkg/db => ./pkg/db
// 	github.com/0B1t322/distanceLearningWebSite/service.auth => ./service.auth
// )

require (
	github.com/0B1t322/distanceLearningWebSite/pkg/db v1.0.0
	github.com/0B1t322/distanceLearningWebSite/service.auth v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)