package micro

import (
	"github.com/MichaelGzy/gooo/go-micro/v1/client"
	"os"
	"os/signal"
	"sync"

	"github.com/MichaelGzy/gooo/go-micro/v1/debug/service/handler"
	"github.com/MichaelGzy/gooo/go-micro/v1/logger"
	"github.com/MichaelGzy/gooo/go-micro/v1/server"
	signalutil "github.com/MichaelGzy/gooo/go-micro/v1/util/signal"
	rtime "runtime"
)

type service struct {
	opts Options

	once sync.Once
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

func (s *service) Init(...Option) {

}
func (s *service) Option() Options {
	return s.Option()
}
func (s *service) Run() error {
	s.opts.Server.Handle(
		s.opts.Server.NewHandler(
			handler.NewHandler(s.opts.Client),
			server.InternalHandler(true),
		))
	if s.opts.Profile != nil {
		rtime.SetMutexProfileFraction(5)
		rtime.SetBlockProfileRate(1)
		if err := s.opts.Profile.Start(); err != nil {
			return err
		}
		defer s.opts.Profile.Stop()
	}
	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Starting [service] %s", s.Name())
	}

	if err := s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	if s.opts.Signal {
		signal.Notify(ch, signalutil.Shutdown()...)
	}

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel
	case <-s.opts.Context.Done():
	}

	return s.Stop()
}
func (s *service) String() string {
	return "micro"
}
func (s *service) Client() client.Client {
	return s.opts.Client
}

func (s *service) Server() server.Server {
	return s.opts.Server
}
func (s *service) Start() error {

}

func (s *service) Stop() error {

}
