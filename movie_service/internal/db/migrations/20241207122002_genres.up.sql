CREATE TABLE public.genres(
      id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      title text NOT NULL UNIQUE,
      created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_genres_title ON public.genres(title);
