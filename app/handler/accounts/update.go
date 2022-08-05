package accounts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

func SaveImage(r *http.Request, imagetype string) (string, error) {
	file, fileHeader, err := r.FormFile(imagetype)
	if err != nil {
		return "", err
	}
	filename := fileHeader.Filename
	defer file.Close()
	filedir := "assets/account/" + imagetype + "/" + filename
	dst, err := os.Create(filedir)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}
	return filename, nil
}

// Handle request for `POST /v1/accounts`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	account := auth.AccountOf(r)
	queries := []string{"display_name", "note", "avatar", "header"}
	for _, v := range queries {
		switch v {
		case "display_name":
			value := r.FormValue(v)
			account.DisplayName = &value
		case "note":
			value := r.FormValue(v)
			account.Note = &value
		case "avatar":
			filename, err := SaveImage(r, "avatar")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			account.Avatar = &filename
		case "header":
			filename, err := SaveImage(r, "header")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			account.Header = &filename
		default:
			httperror.Error(w, http.StatusBadRequest)
			return
		}
	}

	ctx := r.Context()
	if created_user, err := h.app.Dao.Account().UpdateUser(ctx, account); err != nil {
		log.Println(err)
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
