package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	ctx := r.Context()
	if created_user, err := h.app.Dao.Account().CreateUser(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if created_user == nil {
		httperror.Error(w, http.StatusUnauthorized)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(created_user); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}
