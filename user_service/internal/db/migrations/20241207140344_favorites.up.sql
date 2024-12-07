CREATE TABLE public.favorites(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  movie_id int NOT NULL,
  user_id int REFERENCES public.users(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (movie_id, user_id)
);

CREATE INDEX idx_favorites_movie_id ON public.favorites(movie_id);
CREATE INDEX idx_favorites_user_id ON public.favorites(user_id);
