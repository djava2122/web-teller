module gitlab.pactindo.com/ebanking/web-teller

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
// replace github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
// replace github.com/json-iterator/go => github.com/json-iterator/go v1.1.8

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.9
	github.com/micro/go-micro v1.18.0
	github.com/valyala/fastjson v1.6.3

	gitlab.pactindo.com/backend-svc/common v0.0.0-20210525093752-d96a39797155
	gitlab.pactindo.com/ebanking/common v0.0.0-20210523152355-d210512f160b
	gitlab.pactindo.com/ebanking/proto-common v0.0.0-20210525092930-c654da482966
)
