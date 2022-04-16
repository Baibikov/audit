package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
	"google.golang.org/grpc"

	"audit/internal/app/config"
	grpchandler "audit/internal/app/delivery/grpc"
	"audit/internal/app/delivery/grpc/generated"
	"audit/internal/app/repository"
	"audit/internal/app/service"
	"audit/internal/pkg/click"
)

const configPath = "configs/main.yaml"

func app() error {
	out := make(chan bool)
	logrus.Info("initialize application shutdown")
	shutdown(out)

	logrus.Info("initialize config")
	conf, err := config.New(configPath)
	if err != nil {
		return err
	}

	logrus.Info("initialize clickhouse database")
	conn, err := click.Setup(conf.Clickhouse)
	if err != nil {
		return err
	}
	defer func(err error) {
		multierr.AppendInto(&err, conn.Close())
	}(err)

	logrus.Info("initialize storage")
	storage := repository.New(conn)

	logrus.Info("initialize usecases")
	usecase := service.New(storage)

	logrus.Info("initialize grpc server")
	server := grpc.NewServer()
	generated.RegisterAuditLoggerServer(server,  grpchandler.New(usecase))

	logrus.Info("initialize tcp listener")
	listener, err := net.Listen("tcp", ":"+conf.Grpc.Server.Port)
	if err != nil {
		return err
	}

	go func() {
		logrus.Info("serving grpc server")
		if lerr := server.Serve(listener); lerr != nil {
			multierr.AppendInto(&err, lerr)

			return
		}
	}()

	<-out
	logrus.Info("shutdown application")

	return nil
}

func shutdown(out chan<- bool) {
	c := make(chan os.Signal)
	signal.Notify(c,  syscall.SIGINT)

	go func() {
		for {
			select {
			case _, ok := <-c:
				if ok {
					out <- true
					close(c)

					return
				}
			}
		}
	}()
}