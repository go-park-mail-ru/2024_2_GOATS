DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS favorites CASCADE;
DROP TABLE IF EXISTS subscriptions CASCADE;
DROP TYPE IF EXISTS subscription_type_enum;

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

CREATE TYPE public.subscription_type_enum AS ENUM ('pending', 'active');
CREATE TABLE public.subscriptions(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id int REFERENCES public.users(id),
  price int DEFAULT 0,
  status SUBSCRIPTION_TYPE_ENUM DEFAULT 'pending',
  expiration_date timestamp WITH TIME ZONE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_subscriptions_expiration_date ON public.subscriptions(expiration_date);
CREATE INDEX idx_subscriptions_user_id ON public.subscriptions(user_id);
