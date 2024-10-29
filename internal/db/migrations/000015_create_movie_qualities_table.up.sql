CREATE TABLE public.movie_qualities (
    movie_id BIGINT NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
    quality_id BIGINT NOT NULL REFERENCES public.qualities(id) ON DELETE CASCADE,
    video_url TEXT DEFAULT '/static/movies/default.mp4',
    PRIMARY KEY (movie_id, quality_id)
);
CREATE INDEX idx_movie_qualities_movie_id ON public.movie_qualities(movie_id);
CREATE INDEX idx_movie_qualities_quality_id ON public.movie_qualities(quality_id);
