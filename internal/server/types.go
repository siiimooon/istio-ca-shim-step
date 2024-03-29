package server

import (
	"go.uber.org/zap"
	securityapi "istio.io/api/security/v1alpha1"
	"sync"
)

type Server struct {
	securityapi.UnimplementedIstioCertificateServiceServer
	ca            string
	caFingerprint string
	ready         bool
	lock          sync.RWMutex
	logger        *zap.SugaredLogger
}
