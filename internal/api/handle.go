package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/oxxi/accel-one/internal/service"
	"github.com/oxxi/accel-one/pkg/dto"
)

type IContactHandle interface {
	// Interface for handling contact related API requests
	GetContactById(w http.ResponseWriter, r *http.Request)
	UpdateContact(w http.ResponseWriter, r *http.Request)
	DeleteContact(w http.ResponseWriter, r *http.Request)
	SaveContact(w http.ResponseWriter, r *http.Request)
}

type contactHandle struct {
	service service.IContactService // Service used for interacting with contact data
}

// DeleteContact implements IContactHandle.
func (c *contactHandle) DeleteContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}
	err = c.service.Delete(ctx, id)
	if err != nil {
		// Handle potential errors from the service layer
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetContactById implements IContactHandle.
func (c *contactHandle) GetContactById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}
	contact, err := c.service.GetById(ctx, id)
	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Contact not found", http.StatusNotFound)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to encode contact response", http.StatusInternalServerError)
		return
	}
}

// SaveContact implements IContactHandle.
func (c *contactHandle) SaveContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto dto.ContactDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	contact, err := c.service.Save(ctx, dto)
	if err != nil {
		// Handle potential errors from the service layer
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to save contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to encode contact response", http.StatusInternalServerError)
		return
	}
}

// UpdateContact implements IContactHandle.
func (c *contactHandle) UpdateContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	var dto dto.ContactDto
	er := json.NewDecoder(r.Body).Decode(&dto)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	contact, err := c.service.Update(ctx, id, dto)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to encode contact response", http.StatusInternalServerError)
		return
	}
}

func NewContactHandler(s service.IContactService) IContactHandle {
	return &contactHandle{
		service: s,
	}
}
