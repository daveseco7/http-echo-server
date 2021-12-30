package service

import (
	"time"

	"github.com/daveseco7/http-echo-server/internal/config"
)

type service struct {
	wConfig config.WorkloadDurations
}

// NewService instantiates a new Service.
func NewService(workloadConfig config.WorkloadDurations) *service {
	return &service{
		wConfig: workloadConfig,
	}
}

func (s *service) Success() {
	time.Sleep(s.wConfig.SuccessExecDuration)
}

func (s *service) Error400() {
	time.Sleep(s.wConfig.Error400ExecDuration)
}

func (s *service) Error500() {
	time.Sleep(s.wConfig.Error500ExecDuration)
}
