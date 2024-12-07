CREATE TABLE public.movies(
      id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      title text NOT NULL,
      short_description text NOT NULL,
      long_description text NOT NULL,
      card_url text DEFAULT '',
      album_url text DEFAULT '',
      title_url text DEFAULT '',
      release_date date NOT NULL,
      rating decimal(10,2) DEFAULT '0.0',
      movie_type MOVIE_TYPE_ENUM,
      video_url text DEFAULT '',
      country_id int REFERENCES public.countries(id),
      director_id int REFERENCES public.directors(id),
      created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON public.movies(title);
CREATE INDEX idx_movies_movie_type ON public.movies(movie_type);
CREATE INDEX idx_movies_release_date ON public.movies(release_date);
CREATE INDEX idx_movies_country_id ON public.movies(country_id);
