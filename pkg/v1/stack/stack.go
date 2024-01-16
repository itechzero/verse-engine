package stack

import (
	"os"
	"os/signal"
	"sync"

	p "github.com/itechzero/lib-core-go/pkg/v1/provider"

	"github.com/sirupsen/logrus"
)

var runOnce, closeOnce sync.Once

type Stack struct {
	providers []p.Provider
}

func New() *Stack {
	return &Stack{
		providers: make([]p.Provider, 0),
	}
}

func (s *Stack) MustInit(provider p.Provider) {
	name := p.Name(provider)
	if !provider.Enabled() {
		logrus.Warnf("[Stack] Provider [%s] is disabled", name)
		return
	}

	if err := provider.Init(); err != nil {
		logrus.WithError(err).Panicf("[Stack] Provider [%s] is failed to initialize", name)
	}

	s.providers = append(s.providers, provider)
	logrus.Infof("[Stack] Provider [%s] is initialized", name)
}

func (s *Stack) MustRun() {
	runOnce.Do(func() {
		for _, provider := range s.providers {
			if runProvider, ok := provider.(p.RunProvider); ok {
				go s.run(runProvider)
			}
		}
		s.handleInterrupt()
	})
}

func (s *Stack) MustClose() {
	closeOnce.Do(func() {
		for i := len(s.providers) - 1; i >= 0; i-- {
			name := p.Name(s.providers[i])
			logrus.Debugf("[Stack] Provider [%s] is closing", name)

			if err := s.providers[i].Close(); err != nil {
				logrus.WithError(err).Panicf("[Stack] Provider [%s] is failed to close", name)
			}

			logrus.Infof("[Stack] Provider [%s] is closed", name)
		}
	})
}

func (s *Stack) run(provider p.RunProvider) {
	name := p.Name(provider)
	logrus.Debugf("[Stack] Provider [%s] is running", name)

	if err := provider.Run(); err != nil {
		logrus.WithError(err).Panicf("[Stack] Provider [%s] is failed to run", name)
	}
}

// handleInterrupt handles the interrupt signal and calls Close() on the stack.
func (s *Stack) handleInterrupt() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		s.MustClose()
		close(cleanupDone)
	}()
	<-cleanupDone
}
