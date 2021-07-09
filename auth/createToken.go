package auth

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/shaj13/go-guardian/v2/auth"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	user := auth.NewDefaultUser("admin", "1", nil, nil)
	auth.Append(TokenStrategy, token, user)
	body := fmt.Sprintf("token: %s \n", token)
	w.Write([]byte(body))
}
