DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS actors CASCADE;
DROP TABLE IF EXISTS movie_actors CASCADE;
DROP TABLE IF EXISTS collections CASCADE;
DROP TABLE IF EXISTS movie_collections CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS movie_genres CASCADE;
DROP TABLE IF EXISTS countries CASCADE;
DROP TYPE IF EXISTS movie_type_enum;
DROP TYPE IF EXISTS sex_enum;

CREATE TABLE public.genres(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL UNIQUE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_genres_title ON public.genres(title);

CREATE TYPE public.sex_enum AS ENUM ('male', 'female', 'other', 'secret');
CREATE TABLE public.users(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  username text UNIQUE,
  email text NOT NULL UNIQUE,
  avatar_url text,
  password_hash text NOT NULL,
  sex SEX_ENUM,
  birthdate date,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON public.users(email);
CREATE INDEX idx_users_birthdate ON public.users(birthdate);
CREATE INDEX idx_users_created_at ON public.users(created_at);
CREATE INDEX idx_users_updated_at ON public.users(updated_at);

CREATE TABLE public.countries(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  code text NOT NULL,
  flag_url text DEFAULT ''
);

CREATE INDEX idx_countries_title ON public.countries(title);
CREATE INDEX idx_countries_code ON public.countries(code);

CREATE TYPE public.movie_type_enum AS ENUM ('film', 'serial');
CREATE TABLE public.movies(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  description text NOT NULL,
  card_url text DEFAULT '',
  album_url text DEFAULT '',
  release_date date NOT NULL,
  rating decimal(10,2) DEFAULT '0.0',
  movie_type MOVIE_TYPE_ENUM,
  video_url text DEFAULT '',
  country_id int REFERENCES public.countries(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON public.movies(title);
CREATE INDEX idx_movies_movie_type ON public.movies(movie_type);
CREATE INDEX idx_movies_release_date ON public.movies(release_date);
CREATE INDEX idx_movies_country_id ON public.movies(country_id);

CREATE TABLE public.actors(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name text NOT NULL,
  second_name text NOT NULL,
  patronymic text,
  country_id int REFERENCES public.countries(id),
  photo_url text DEFAULT '',
  birthdate date,
  biography text DEFAULT ''
);

CREATE INDEX idx_actors_country_id ON public.actors(country_id);

CREATE TABLE public.movie_actors(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  movie_id int REFERENCES public.movies(id),
  actor_id int REFERENCES public.actors(id)
);

CREATE INDEX idx_movie_actors_movie_id ON public.movie_actors(movie_id);
CREATE INDEX idx_movie_actors_actor_id ON public.movie_actors(actor_id);

CREATE TABLE public.movie_genres(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  movie_id int REFERENCES public.movies(id),
  genre_id int REFERENCES public.genres(id)
);

CREATE INDEX idx_movie_genres_movie_id ON public.movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON public.movie_genres(genre_id);

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

CREATE TABLE public.movie_collections(
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  movie_id int REFERENCES public.movies(id),
  collection_id int REFERENCES public.collections(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movie_collections_movie_id ON public.movie_collections(movie_id);
CREATE INDEX idx_movie_collections_collection_id ON public.movie_collections(collection_id);
