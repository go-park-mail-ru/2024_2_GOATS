CREATE TABLE public.movie_collections(
     id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
     movie_id int REFERENCES public.movies(id),
     collection_id int REFERENCES public.collections(id),
     created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movie_collections_movie_id ON public.movie_collections(movie_id);
CREATE INDEX idx_movie_collections_collection_id ON public.movie_collections(collection_id);
