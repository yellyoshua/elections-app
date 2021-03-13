package models

// BodyLoginUser body params of login request
type BodyLoginUser struct {
	Identifier string `json:"identifier" validate:"nonzero"`
	Password   string `json:"password" validate:"nonzero"`
}
