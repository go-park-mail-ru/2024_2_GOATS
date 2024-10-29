CREATE TABLE public.collections(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title TEXT NOT NULL UNIQUE,
  card_url TEXT DEFAULT '/static/collections/default_card.png',
  album_url TEXT DEFAULT '/static/collections/default_poster.png',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_collections_title ON public.collections(title);
CREATE INDEX idx_collections_created_at ON public.collections(created_at);
