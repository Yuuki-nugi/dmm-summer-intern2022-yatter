package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `POST /v1/accounts`
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	max_id := r.FormValue("max_id")
	since_id := r.FormValue("since_id")
	limit := r.FormValue("limit")

	ctx := r.Context()
	if statuses, err := h.app.Dao.Timelines().GetPublicTimelines(ctx, max_id, since_id, limit); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if statuses == nil {
		httperror.Error(w, http.StatusUnauthorized)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statuses); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}
