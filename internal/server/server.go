package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HuaiyuZhang/benben-ml/internal/config"
	"github.com/HuaiyuZhang/benben-ml/internal/model"
	"github.com/gorilla/mux"
)

// Server represents our HTTP server
type Server struct {
	Router *mux.Router
	Model  model.Model
	Config *config.Config
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Model:  model.NewModel(),
		Config: cfg,
	}

	// Load the model
	if err := s.Model.Load(cfg.ModelPath); err != nil {
		log.Printf("Warning: Failed to load model: %v", err)
	}

	return s
}

// InitializeRoutes sets up all the routes for our server
func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/predict", s.handlePredict).Methods("POST")
	s.Router.HandleFunc("/health", s.handleHealth).Methods("GET")
}

// handlePredict handles prediction requests
func (s *Server) handlePredict(w http.ResponseWriter, r *http.Request) {
	var req model.PredictRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	prediction, err := s.Model.Predict(req.Features)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.PredictResponse{
		Prediction: prediction,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleHealth handles health check requests
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
