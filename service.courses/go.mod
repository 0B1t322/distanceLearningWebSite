module github.com/0B1t322/service.courses

go 1.15

replace (
	github.com/0B1t322/distanceLearningWebSite => ../
	github.com/0B1t322/distanceLearningWebSite/service.auth => ../service.auth
)

require (
	github.com/0B1t322/distanceLearningWebSite v0.0.0
	github.com/0B1t322/distanceLearningWebSite/service.auth v0.0.0
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/jmoiron/sqlx v1.3.1 // indirect
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.4 // indirect
	gorm.io/gorm v1.20.12
)
