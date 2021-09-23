package main

import (
	"strconv"
	"time"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/micro"
	"gitlab.pactindo.com/ebanking/common/pg"
	"gitlab.pactindo.com/ebanking/common/redis"

	"gitlab.pactindo.com/ebanking/web-teller/service"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func main() {
	// Define service
	svc := micro.NewService(
		micro.ServiceName("web-teller"),
		micro.Config(
			"DB_URL",
			"REDIS_URL",
			"REDIS_POOLSIZE",
		),
		micro.RequestTimeout(time.Second*58),
	)

	log.InfoS("starting service: " + micro.GetServiceName())

	conf := micro.GetConfig()

	poolSize, _ := strconv.Atoi(conf["REDIS_POOLSIZE"])

	redis.Init(conf["REDIS_URL"], poolSize)

	dbURL := conf["DB_URL"]

	log.InfoS("DB_URL: " + dbURL)

	pg.Init(conf["DB_URL"], 0)

	service.Init(svc.Client())

	wtproto.RegisterWebTellerHandler(svc.Server(), new(service.WebTellerHandler))

	// Run Service
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
