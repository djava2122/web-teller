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
	github.com/nats-io/nats.go v1.9.2 // indirect
	github.com/valyala/fastjson v1.6.3
	gitlab.pactindo.com/ebanking/common v0.0.0-20210523152355-d210512f160b
	gitlab.pactindo.com/ebanking/proto-common v0.0.0-20210525092930-c654da482966
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
