package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"
)

type Queries struct {
	MaxId   int
	SinceId int
	Limit   int
}

// Handle request for `POST /v1/accounts`
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	var formValues Queries
	queries := []string{"max_id", "since_id", "limit"}
	for _, v := range queries {
		value, err := strconv.Atoi(r.FormValue(v))
		if err != nil {
			httperror.Error(w, http.StatusBadRequest)
			return
		}
		switch v {
		case "max_id":
			formValues.MaxId = value
		case "since_id":
			formValues.SinceId = value
		case "limit":
			formValues.Limit = value
		default:
			httperror.Error(w, http.StatusBadRequest)
			return
		}
	}

	ctx := r.Context()
	if statuses, err := h.app.Dao.Timelines().GetPublicTimelines(ctx, formValues.MaxId, formValues.SinceId, formValues.Limit); err != nil {
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
