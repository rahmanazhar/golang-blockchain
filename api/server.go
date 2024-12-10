package api

import (
	"blockchain/pkg/blockchain"
	"encoding/json"
	"net/http"
	"os"
)

type BlockchainServer struct {
	bc *blockchain.Blockchain
}

type BlockData struct {
	Data string `json:"data"`
}

type BlockchainResponse struct {
	Blocks []Block `json:"blocks"`
	Valid  bool    `json:"valid"`
}

type Block struct {
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
	Nonce     int    `json:"nonce"`
}

func NewBlockchainServer() *BlockchainServer {
	return &BlockchainServer{
		bc: blockchain.NewBlockchain(),
	}
}

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		corsOrigin := os.Getenv("CORS_ORIGIN")
		if corsOrigin == "" {
			corsOrigin = "http://localhost:3000"
		}
		
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func (s *BlockchainServer) HandleAddBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var blockData BlockData
	if err := json.NewDecoder(r.Body).Decode(&blockData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.bc.AddBlock(blockData.Data)
	w.WriteHeader(http.StatusCreated)
}

func (s *BlockchainServer) HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var blocks []Block
	for _, b := range s.bc.Blocks {
		blocks = append(blocks, Block{
			Timestamp: b.Timestamp,
			Data:      b.Data,
			Hash:      b.Hash,
			PrevHash:  b.PrevHash,
			Nonce:     b.Nonce,
		})
	}

	response := BlockchainResponse{
		Blocks: blocks,
		Valid:  s.bc.IsValid(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *BlockchainServer) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/blockchain", enableCORS(s.HandleGetBlockchain))
	mux.HandleFunc("/block", enableCORS(s.HandleAddBlock))
	return mux
}