CREATE TABLE public.users(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  username text UNIQUE,
  email text NOT NULL UNIQUE,
  avatar_url text DEFAULT '/static/user_avatars/default.jpg',
  password_hash text NOT NULL,
  birthdate date,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON public.users(email);
CREATE INDEX idx_users_birthdate ON public.users(birthdate);
CREATE INDEX idx_users_created_at ON public.users(created_at);
CREATE INDEX idx_users_updated_at ON public.users(updated_at);
