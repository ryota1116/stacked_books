package handler

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"net/http"
	"strconv"
	"time"
)

func setUserSession(w http.ResponseWriter, user model.User) {


	fmt.Println(user.Id)
	expiration := time.Now()
	expiration.AddDate(0, 0, 7)
	cookie := http.Cookie{
		Name:       "user_id",
		Value:      strconv.Itoa(user.Id),
		Expires:    expiration,
	}

	//cookie := http.Cookie{
	//	Name:       "user_session_key",
	//	Value:      uuid.Generate(uuid.Bits),
	//	Expires:    expiration,
	//}


	http.SetCookie(w, &cookie)
}

func readUserSession(w http.ResponseWriter, r *http.Request)  {

}