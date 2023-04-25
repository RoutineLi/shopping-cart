package emqx_api

type CreateAuthUserRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type CreateAuthUserResponse struct {
	IsSuperUser bool   `json:"is_superuser"`
	UserId      string `json:"user_id"`
}
