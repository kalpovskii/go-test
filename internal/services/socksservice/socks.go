package socksservice

import (
	"awesomeProject/internal/repositories/socksrepository"
	"github.com/gorilla/mux"
	"net/http"
)

type SocksService struct {
	socksRepository socksrepository.SocksRepository
}

func New(socksRepository socksrepository.SocksRepository) *SocksService {
	return &SocksService{
		socksRepository: socksRepository,
	}
}

func (ss *SocksService) GetHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/socks", ss.Create).Methods(http.MethodGet)
	router.HandleFunc("/api/socks", ss.Create).Methods(http.MethodGet)
}
