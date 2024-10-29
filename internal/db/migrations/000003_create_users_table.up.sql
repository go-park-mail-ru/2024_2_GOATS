CREATE TABLE public.users(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  username TEXT UNIQUE,
  email TEXT NOT NULL UNIQUE,
  avatar_url TEXT DEFAULT '/static/avatars/default.png',
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON public.users(email);
CREATE INDEX idx_users_created_at ON public.users(created_at);
CREATE INDEX idx_users_updated_at ON public.users(updated_at);
