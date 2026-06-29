package domain

type LoginResult struct {
	Token     Token  `json:"token"`
	UserType  int    `json:"user_type"`
	Approved  bool   `json:"approved"`
}
