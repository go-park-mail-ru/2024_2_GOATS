CREATE TABLE public.collections(
       id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       title text NOT NULL UNIQUE,
       card_url text,
       album_url text,
       conditions json,
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_collections_title ON public.collections(title);
CREATE INDEX idx_collections_created_at ON public.collections(created_at);
