package models

type MatchActorPhrasePrefix struct {
	FullName string `json:"full_name"`
}

type ActorQuery struct {
	MatchActorPhrasePrefix MatchActorPhrasePrefix `json:"match_phrase_prefix"`
}

type SearchActorQuery struct {
	ActorQuery ActorQuery `json:"query"`
}

type ActorSource struct {
	ID          string `json:"id"`
	Name        string `json:"full_name"`
	PhotoBigURL string `json:"photo_big_url"`
}

type ActorHit struct {
	ActorSource ActorSource `json:"_source"`
}

type ActorHits struct {
	ActorHits []ActorHit `json:"hits"`
}

type ActorESResponse struct {
	ActorHits ActorHits `json:"hits"`
}

type MatchMoviePhrasePrefix struct {
	Title string `json:"title"`
}

type MovieQuery struct {
	MatchMoviePhrasePrefix MatchMoviePhrasePrefix `json:"match_phrase_prefix"`
}

type SearchMovieQuery struct {
	MovieQuery MovieQuery `json:"query"`
}

type MovieSource struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Rating   float32 `json:"rating"`
	AlbumURL string  `json:"album_url"`
	CardURL  string  `json:"card_url"`
}

type MovieHit struct {
	MovieSource MovieSource `json:"_source"`
}

type MovieHits struct {
	MovieHits []MovieHit `json:"hits"`
}

type MovieESResponse struct {
	MovieHits MovieHits `json:"hits"`
}
