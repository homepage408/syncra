package model

type Post struct {
	ID       string
	AuthorID string
	Title    string
	Body     string
}

type AuthPayload struct {
	AccessToken  string
	RefreshToken string
}
