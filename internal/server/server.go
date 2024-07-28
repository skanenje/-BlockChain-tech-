package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"swapBlock/internal/agreement"
	"swapBlock/internal/node"
)

type Server struct {
	node *node.Node
}

func NewServer(nodeID string) *Server {
	return &Server{
		node: node.NewNode(nodeID),
	}
}

func (s *Server) HandleHome(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the templates directory
	templatesDir, err := filepath.Abs("web/templates")
	if err != nil {
		http.Error(w, "Unable to find templates", http.StatusInternalServerError)
		return
	}

	// Parse the template file
	tmpl, err := template.ParseFiles(filepath.Join(templatesDir, "index.html"))
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	// Collect all agreements from the blockchain
	var agreements []agreement.SaleAgreement
	for _, block := range s.node.Blockchain.Chain {
		if block.Index > 0 { // Skip genesis block
			agreements = append(agreements, block.Agreement)
		}
	}

	// Execute the template with the agreements data
	err = tmpl.Execute(w, agreements)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, _ := strconv.Atoi(r.FormValue("purchasePrice"))

	newAgreement := agreement.NewSaleAgreement(
		r.FormValue("seller"),
		r.FormValue("buyer"),
		r.FormValue("vehicleMake"),
		r.FormValue("registrationNumber"),
		price,
	)

	s.node.AddAgreement(newAgreement)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) HandleBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.node.Blockchain)
}

func (s *Server) Start() {
	http.HandleFunc("/", s.HandleHome)
	http.HandleFunc("/create", s.HandleCreate)
	http.HandleFunc("/blockchain", s.HandleBlockchain)

	http.ListenAndServe(":8080", nil)
}
