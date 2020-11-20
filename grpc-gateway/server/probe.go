package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"grpc-gateway/server/healthcheck"
)

type HealthServer struct {
	healthcheck.UnimplementedHealthServer
}

func (s *HealthServer) Check(context.Context, *healthcheck.HealthCheckRequest) (*healthcheck.HealthCheckResponse, error) {
	logrus.Info("Serving the Check request for health check")
	return &healthcheck.HealthCheckResponse{
		Status: healthcheck.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthServer) Watch(req *healthcheck.HealthCheckRequest, server healthcheck.Health_WatchServer) error {
	logrus.Info("Serving the Watch request for health check")
	return server.Send(&healthcheck.HealthCheckResponse{
		Status: healthcheck.HealthCheckResponse_SERVING,
	})
}

func NewHealthChecker() *HealthServer {
	return &HealthServer{}
}
