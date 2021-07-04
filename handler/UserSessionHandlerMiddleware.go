package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model"
	"net/http"
	"time"
)

func setUserSession(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	expiration := time.Now()
	expiration.AddDate(0, 0, 180)
	cookie := http.Cookie{
		Name:       "user_id",
		Value:      "",
		Expires:    expiration,
	}

	http.SetCookie(w, &cookie)
}

func readUserSession(w http.ResponseWriter, r *http.Request)  {

}