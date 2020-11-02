package autoops

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	AdmissionServerCertFile = "/autoops-data/admission-server/tls.crt"
	AdmissionServerKeyFile  = "/autoops-data/admission-server/tls.key"
)

func ListenAndServeAdmission(s *http.Server) (err error) {
	log.Println("listening at :443")
	return s.ListenAndServeTLS(AdmissionServerCertFile, AdmissionServerKeyFile)
}

func RunAdmissionServer(s *http.Server) (err error) {
	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		chErr <- ListenAndServeAdmission(s)
	}()

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Println("signal caught:", sig.String())
		_ = s.Shutdown(context.Background())
	}

	return
}
