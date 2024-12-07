// HTTP API Handler
package http

import (
    "encoding/json"
    "net/http"
    "github.com/imanimen/ddd-pr/internal/app"
)

type UserHandler struct {
    appService *app.UserAppService
}

func NewUserHandler(appService *app.UserAppService) *UserHandler {
    return &UserHandler{appService: appService}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    user, err := h.appService.RegisterUser(req.Name, req.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}
