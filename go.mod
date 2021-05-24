module gitlab.pactindo.com/ebanking/web-teller

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/json-iterator/go v1.1.9
	github.com/micro/go-micro v1.18.0

	gitlab.pactindo.com/backend-svc/common v0.0.0-20210409083619-b2131b449767
)
