module git.pactindo.com/ebanking/web-teller

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0

// replace github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
// replace github.com/json-iterator/go => github.com/json-iterator/go v1.1.8

require (
	git.pactindo.com/ebanking/common v0.0.0-20220608025919-c8b7157b1bda
	git.pactindo.com/ebanking/proto-common v0.0.0-20220607040155-6d18724f84e6
	git.pactindo.com/ebanking/proto-ibmb v1.0.2-0.20220607040307-2fe200d58b61
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.9
	github.com/lib/pq v1.3.0 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.27 // indirect
	github.com/nats-io/nats-server/v2 v2.1.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/valyala/fastjson v1.6.3
	go.etcd.io/bbolt v1.3.4 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20191216173652-a0e659d51361 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
