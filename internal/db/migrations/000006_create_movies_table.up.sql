CREATE TABLE public.movies(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title TEXT NOT NULL,
  short_description TEXT NOT NULL,
  long_description TEXT NOT NULL,
  card_url TEXT DEFAULT '/static/movies/default_card.png',
  album_url TEXT DEFAULT '/static/movies/default_poster.png',
  title_url TEXT DEFAULT '/static/movies/default_title.png',
  release_date DATE NOT NULL,
  movie_type MOVIE_TYPE_ENUM NOT NULL,
  country_id BIGINT NOT NULL REFERENCES public.countries(id),
  director_id BIGINT NOT NULL REFERENCES public.directors(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON public.movies(title text_pattern_ops);
CREATE INDEX idx_movies_movie_type ON public.movies(movie_type);
CREATE INDEX idx_movies_release_date ON public.movies(release_date);
CREATE INDEX idx_movies_director_id ON public.movies(director_id);
CREATE INDEX idx_movies_country_id ON public.movies(country_id);
