package dto

type DBRegisterData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type DBUser struct {
	ID        int
	Email     string
	Username  string
	Password  string
	AvatarURL string
}

type DBFavorite struct {
	UserID  int
	MovieID int
}

type DBMovieShortInfo struct {
	ID          int
	Title       string
	CardURL     string
	AlbumURL    string
	Rating      float32
	ReleaseDate string
	MovieType   string
	Country     string
}
