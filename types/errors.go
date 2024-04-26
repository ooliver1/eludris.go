package types

import "errors"

var (
	ErrNoConfig                = errors.New("config is required")
	ErrNoToken                 = errors.New("config.Token is required")
	ErrNoApiUrl                = errors.New("config.ApiUrl is required")
	ErrGatewayAlreadyConnected = errors.New("gateway is already connected")
	ErrGatewayNotConnected     = errors.New("not connected ŧo the gateway")
	ErrNoGateway               = errors.New("no gateway configured")
)
