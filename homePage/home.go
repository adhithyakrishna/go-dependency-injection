package homePage

import (
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Inside home")
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("from home page"))
}
