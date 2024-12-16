package models

// MatchActorPhrasePrefix ES struct
type MatchActorPhrasePrefix struct {
	FullName string `json:"full_name"`
}

// ActorQuery ES struct
type ActorQuery struct {
	MatchActorPhrasePrefix MatchActorPhrasePrefix `json:"match_phrase_prefix"`
}

// SearchActorQuery ES struct
type SearchActorQuery struct {
	ActorQuery ActorQuery `json:"query"`
}

// ActorSource ES struct
type ActorSource struct {
	ID          string `json:"id"`
	Name        string `json:"full_name"`
	PhotoBigURL string `json:"photo_big_url"`
}

// ActorHit ES struct
type ActorHit struct {
	ActorSource ActorSource `json:"_source"`
}

// ActorHits ES struct
type ActorHits struct {
	ActorHits []ActorHit `json:"hits"`
}

// ActorESResponse ES struct
type ActorESResponse struct {
	ActorHits ActorHits `json:"hits"`
}

// MatchMoviePhrasePrefix ES struct
type MatchMoviePhrasePrefix struct {
	Title string `json:"title"`
}

// MovieQuery ES struct
type MovieQuery struct {
	MatchMoviePhrasePrefix MatchMoviePhrasePrefix `json:"match_phrase_prefix"`
}

// SearchMovieQuery ES struct
type SearchMovieQuery struct {
	MovieQuery MovieQuery `json:"query"`
}

// MovieSource ES struct
type MovieSource struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Rating   float32 `json:"rating"`
	AlbumURL string  `json:"album_url"`
	CardURL  string  `json:"card_url"`
}

// MovieHit ES struct
type MovieHit struct {
	MovieSource MovieSource `json:"_source"`
}

// MovieHits ES struct
type MovieHits struct {
	MovieHits []MovieHit `json:"hits"`
}

// MovieESResponse ES struct
type MovieESResponse struct {
	MovieHits MovieHits `json:"hits"`
}
