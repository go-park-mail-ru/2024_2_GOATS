CREATE TABLE public.movie_genres(
  movie_id BIGINT NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
  genre_id BIGINT NOT NULL REFERENCES public.genres(id) ON DELETE CASCADE,
  PRIMARY KEY (movie_id, genre_id)
);

CREATE INDEX idx_movie_genres_movie_id ON public.movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON public.movie_genres(genre_id);
