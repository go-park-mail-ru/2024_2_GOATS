CREATE TABLE public.movie_genres(
        id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        movie_id int REFERENCES public.movies(id),
        genre_id int REFERENCES public.genres(id)
);

CREATE INDEX idx_movie_genres_movie_id ON public.movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON public.movie_genres(genre_id);
