package httpserver

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
	"transaction-service/config"
)

const (
	_defaultReadTimeOut 	= 5 * time.Second
	_defaultWriteTimeOut 	= 5 * time.Second
	_defaultAddr  			= ":80"
	_defaultShutdownTimeOut = 3 * time.Second
)

type Server struct {
	server 			*http.Server
	notify 			chan error
	shutDowntimeOut time.Duration
	cfg			 	*config.Config
}

func New(handler http.Handler, cfg *config.Config, opts ...Option) *Server {
	httpServer := &http.Server{}

	if (cfg.App.Environment == "development" || cfg.App.Environment == "testing") && cfg.HTTPServer.UseSSL {
		
		cfgClient := &tls.Config{
			MinVersion: tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		httpServer.Addr = ":" + cfg.HTTPServer.Port
		httpServer.Handler = handler
		httpServer.ReadTimeout = _defaultReadTimeOut
		httpServer.WriteTimeout = _defaultWriteTimeOut
		httpServer.TLSConfig = cfgClient
		httpServer.TLSNextProto =  make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

	} else {
		httpServer.Addr = ":" + cfg.HTTPServer.Port
		httpServer.Handler = handler
		httpServer.ReadTimeout = _defaultReadTimeOut
		httpServer.WriteTimeout = _defaultWriteTimeOut
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
		shutDowntimeOut: _defaultShutdownTimeOut,
		cfg: cfg,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		if s.cfg.HTTPServer.UseSSL {
			s.notify <- s.server.ListenAndServeTLS(s.cfg.BaseDir+s.cfg.HTTPServer.SSLCert, s.cfg.BaseDir+s.cfg.HTTPServer.SSLKey)
		} else {
			s.notify <- s.server.ListenAndServe()
		}
		close(s.notify)
	}()
}

func (s * Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutDowntimeOut)
	defer cancel()

	return s.server.Shutdown(ctx)
}