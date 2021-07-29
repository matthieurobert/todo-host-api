package auth

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shaj13/go-guardian/v2/auth"

	"github.com/todo-host/todo-host-api/entity"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	_, userInfo, _ := Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())
	defaultUser := auth.NewDefaultUser(user.Username, string(rune(user.Id)), nil, nil)
	auth.Append(TokenStrategy, token, defaultUser)
	body := fmt.Sprintf("token: %s \n", token)
	w.Write([]byte(body))
}
