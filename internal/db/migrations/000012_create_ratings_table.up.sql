CREATE TABLE public.ratings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    movie_id BIGINT NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
    episode_id BIGINT REFERENCES public.episodes(id) ON DELETE CASCADE,
    rating DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_id, movie_id, episode_id)
);
CREATE INDEX idx_ratings_user_id ON public.ratings(user_id);
CREATE INDEX idx_ratings_movie_id ON public.ratings(movie_id);
CREATE INDEX idx_ratings_episode_id ON public.ratings(episode_id);
CREATE UNIQUE INDEX idx_ratings_user_movie_episode ON public.ratings(user_id, movie_id, episode_id);
