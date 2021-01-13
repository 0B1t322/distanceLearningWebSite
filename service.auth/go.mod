module github.com/0B1t322/service.auth

go 1.15

replace github.com/0B1t322/distanceLearningWebSite => ../

require (
	github.com/0B1t322/distanceLearningWebSite v0.0.0
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20210113131315-ba0562f347e0 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210113155445-facbc42f5e06 // indirect
	google.golang.org/grpc v1.34.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gorm.io/gorm v1.20.11
)
