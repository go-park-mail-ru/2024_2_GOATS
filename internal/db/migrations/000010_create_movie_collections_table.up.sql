CREATE TABLE public.movie_collections(
  movie_id BIGINT NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
  collection_id BIGINT NOT NULL REFERENCES public.collections(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (movie_id, collection_id)
);

CREATE INDEX idx_movie_collections_movie_id ON public.movie_collections(movie_id);
CREATE INDEX idx_movie_collections_collection_id ON public.movie_collections(collection_id);
