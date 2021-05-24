module gitlab.pactindo.com/ebanking/web-teller

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.1
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.9
	github.com/lib/pq v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.4.0

	gitlab.pactindo.com/ebanking/common v0.0.0-20210503093142-af90e4018adc
	gitlab.pactindo.com/ebanking/proto-common v0.0.0-20210410071444-7937aec6b8c4
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)
