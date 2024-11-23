package dto

type RepoRegisterData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type RepoUser struct {
	ID        int
	Email     string
	Username  string
	Password  string
	AvatarURL string
}

type RepoFavorite struct {
	UserID  int
	MovieID int
}

type RepoMovieShortInfo struct {
	ID          int
	Title       string
	CardURL     string
	AlbumURL    string
	Rating      float32
	ReleaseDate string
	MovieType   string
	Country     string
}
