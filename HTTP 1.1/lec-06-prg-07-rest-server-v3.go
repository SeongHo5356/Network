package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type MembershipHandler struct {
	database map[string]string
}

func MakeMembership() *MembershipHandler {
	return &MembershipHandler{
		database: make(map[string]string),
	}
}

// POST
func (m *MembershipHandler) create(id, value string) map[string]string {
	if _, exists := m.database[id]; exists {
		return map[string]string{id: "None"}
	}
	m.database[id] = value
	return map[string]string{id: m.database[id]}
}

// GET
func (m *MembershipHandler) read(id string) map[string]string {
	if value, exists := m.database[id]; exists {
		return map[string]string{id: value}
	}
	return map[string]string{id: "None"}
}

// PUT
func (m *MembershipHandler) update(id, value string) map[string]string {
	if _, exists := m.database[id]; exists {
		return map[string]string{id: value}
	}
	m.database[id] = value
	return map[string]string{id: "None"}
}

// DELETE
func (m *MembershipHandler) delete(id string) map[string]string {
	if _, exists := m.database[id]; exists {
		delete(m.database, id)
		return map[string]string{id: "removed"}
	}
	return map[string]string{id: "None"}
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

var myManager = MakeMembership()

func main() {
	http.HandleFunc("/membership_api/", func(w http.ResponseWriter, r *http.Request) {
		member_id := strings.TrimPrefix(r.URL.Path, "/membership_api/")
		if member_id == "" {
			http.Error(w, "MemberID not provided", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPost:
			result := myManager.create(member_id, r.FormValue(member_id))
			respondJSON(w, result)
		case http.MethodGet:
			result := myManager.read(member_id)
			respondJSON(w, result)
		case http.MethodPut:
			result := myManager.update(member_id, r.FormValue(member_id))
			respondJSON(w, result)
		case http.MethodDelete:
			result := myManager.delete(member_id)
			respondJSON(w, result)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("HTTP server started at http://localhost:5000")

	http.ListenAndServe(":5000", nil)
}
