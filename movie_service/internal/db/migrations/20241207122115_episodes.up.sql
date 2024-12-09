CREATE TABLE public.episodes (
         id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
         season_id INT NOT NULL REFERENCES public.seasons(id) ON DELETE CASCADE,
         episode_number INT NOT NULL,
         title TEXT NOT NULL,
         description TEXT NOT NULL,
         rating DECIMAL(10, 2) NOT NULL DEFAULT '0.0',
         preview_url TEXT DEFAULT '/static/serials/default.png',
         video_url TEXT DEFAULT '/static/serials/default.mp4',
         release_date DATE NOT NULL,
         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

         UNIQUE (season_id, episode_number)
);

CREATE INDEX idx_episodes_season_id ON public.episodes(season_id);
