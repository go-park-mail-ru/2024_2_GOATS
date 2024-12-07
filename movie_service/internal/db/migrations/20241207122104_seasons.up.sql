CREATE TABLE public.seasons (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        movie_id BIGINT NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
        season_number INT NOT NULL,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
        rating DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
        preview_url TEXT DEFAULT '/static/serials/default.png',
        release_date DATE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

        UNIQUE (movie_id, season_number)
);

CREATE INDEX idx_seasons_movie_id ON public.seasons(movie_id);
