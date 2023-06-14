package socksservice

import (
	"awesomeProject/internal/repositories/socksrepository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	router.HandleFunc("/api/socks/income", ss.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/socks", ss.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/socks/outcome", ss.Delete).Methods(http.MethodPost)

	// CORS
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	method := handlers.AllowedMethods([]string{"POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(header, method, origins)(router)
}

type CreateRequest struct {
	Color      string `json:"color"`
	CottonPart int64  `json:"cottonPart"`
	Quantity   int64  `json:"quantity"`
}

func (ss *SocksService) Create(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s := socksrepository.Socks{
		Color:      req.Color,
		CottonPart: req.CottonPart,
		Quantity:   req.Quantity,
	}
	if err := ss.socksRepository.Create(r.Context(), &s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ss *SocksService) Get(w http.ResponseWriter, r *http.Request) {
	color := r.URL.Query().Get("color")
	operation := r.URL.Query().Get("operation")
	cottonPartStr := r.URL.Query().Get("cottonPart")

	cottonPart, err := strconv.Atoi(cottonPartStr)
	if err != nil {
		fmt.Printf("error while conv to str: %s", err)
	}
	s := socksrepository.Socks{
		Color:      color,
		CottonPart: int64(cottonPart),
	}

	totalCount, err := ss.socksRepository.Get(r.Context(), &s, operation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "%d", totalCount)
	if err != nil {
		return
	}
}

func (ss *SocksService) Delete(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s := socksrepository.Socks{
		Color:      req.Color,
		CottonPart: req.CottonPart,
		Quantity:   req.Quantity,
	}

	if err := ss.socksRepository.Delete(r.Context(), &s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
