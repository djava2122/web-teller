package main

import (
	"strconv"
	"time"

	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/micro"
	"git.pactindo.com/ebanking/common/pg"
	"git.pactindo.com/ebanking/common/redis"

	"git.pactindo.com/ebanking/web-teller/service"

	wtproto "git.pactindo.com/ebanking/web-teller/proto"
)

func main() {
	// Define service //
	svc := micro.NewService(
		micro.ServiceName("web-teller"),
		micro.Config(
			"DB_URL",
			"REDIS_URL",
			"REDIS_POOLSIZE",
			"URL_SIKOPATUH",
			"URL_MGATE",
			"USER_SIKOPATUH",
			"PASS_SIKOPATUH",
			"FILE_LOCATION",
		),
		micro.RequestTimeout(time.Second*137),
	)

	log.InfoS("starting service: " + micro.GetServiceName())

	conf := micro.GetConfig()

	poolSize, _ := strconv.Atoi(conf["REDIS_POOLSIZE"])

	redis.Init(conf["REDIS_URL"], poolSize)

	dbURL := conf["DB_URL"]

	log.InfoS("DB_URL: " + dbURL)
	log.InfoS("FILE_LOCATION: " + conf["FILE_LOCATION"])

	pg.Init(conf["DB_URL"], 0)

	service.Init(svc.Client())

	wtproto.RegisterWebTellerHandler(svc.Server(), new(service.WebTellerHandler))

	// Run Service
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
