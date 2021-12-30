package config

import (
	"time"
)

const (
	httpServerReadTimeout  = 60
	httpServerWriteTimeout = 10
	httpServerIdleTimeout  = 120

	WorkloadDurationsSuccessExecDuration  = 1
	WorkloadDurationsError400ExecDuration = 3
	WorkloadDurationsError500ExecDuration = 5
)

// Config represents the configs needed by the service.
type Config struct {
	Server    *HTTPServer        `yaml:"server"`
	Workloads *WorkloadDurations `yaml:"workload"`
}

// HTTPServer represents an HTTP server configuration.
type HTTPServer struct {
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

type WorkloadDurations struct {
	SuccessExecDuration  time.Duration `yaml:"successExecDuration"`
	Error400ExecDuration time.Duration `yaml:"error400ExecDuration"`
	Error500ExecDuration time.Duration `yaml:"error500ExecDuration"`
}

func (c *Config) HTTPServerConfig() HTTPServer {
	if c.Server == nil {
		return HTTPServer{
			ReadTimeout:  httpServerReadTimeout,
			WriteTimeout: httpServerWriteTimeout,
			IdleTimeout:  httpServerIdleTimeout,
		}
	}

	return *c.Server
}

func (c *Config) WorkloadConfig() WorkloadDurations {
	if c.Workloads == nil {
		return WorkloadDurations{
			SuccessExecDuration:  WorkloadDurationsSuccessExecDuration,
			Error400ExecDuration: WorkloadDurationsError400ExecDuration,
			Error500ExecDuration: WorkloadDurationsError500ExecDuration,
		}
	}

	return *c.Workloads
}
