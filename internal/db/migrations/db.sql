DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS actors CASCADE;
DROP TABLE IF EXISTS directors CASCADE;
DROP TABLE IF EXISTS movie_actors CASCADE;
DROP TABLE IF EXISTS collections CASCADE;
DROP TABLE IF EXISTS movie_collections CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS movie_genres CASCADE;
DROP TABLE IF EXISTS countries CASCADE;
DROP TYPE IF EXISTS movie_type_enum;

CREATE TABLE public.genres(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL UNIQUE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_genres_title ON public.genres(title);

CREATE TABLE public.directors(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name text NOT NULL,
  second_name text NOT NULL
);

CREATE INDEX idx_directors_name_pattern ON public.directors (first_name text_pattern_ops, second_name text_pattern_ops);

CREATE TABLE public.users(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  username text UNIQUE,
  email text NOT NULL UNIQUE,
  avatar_url text DEFAULT '/static/avatars/default.png',
  password_hash text NOT NULL,
  birthdate date,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON public.users(email);
CREATE INDEX idx_users_birthdate ON public.users(birthdate);
CREATE INDEX idx_users_created_at ON public.users(created_at);
CREATE INDEX idx_users_updated_at ON public.users(updated_at);

CREATE TABLE public.countries(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  code text NOT NULL,
  flag_url text DEFAULT '/static/flags/default.png'
);

CREATE INDEX idx_countries_title ON public.countries(title);
CREATE INDEX idx_countries_code ON public.countries(code);

CREATE TYPE public.movie_type_enum AS ENUM ('film', 'serial');
CREATE TABLE public.movies(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL,
  short_description text NOT NULL,
  long_description text NOT NULL,
  card_url text DEFAULT '/static/movies/default_card.png',
  album_url text DEFAULT '/static/movies/default_poster.png',
  title_url text DEFAULT '/static/movies/default_title.png',
  release_date date NOT NULL,
  rating decimal(10,2) DEFAULT '0.0',
  movie_type MOVIE_TYPE_ENUM NOT NULL,
  country_id bigint REFERENCES public.countries(id),
  director_id bigint REFERENCES public.directors(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON public.movies(title text_pattern_ops);
CREATE INDEX idx_movies_movie_type ON public.movies(movie_type);
CREATE INDEX idx_movies_release_date ON public.movies(release_date);
CREATE INDEX idx_movies_rating ON public.movies(rating);
CREATE INDEX idx_movies_director_id ON public.movies(director_id);
CREATE INDEX idx_movies_country_id ON public.movies(country_id);

CREATE TABLE public.actors(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name text NOT NULL,
  second_name text NOT NULL,
  country_id bigint REFERENCES public.countries(id),
  small_photo_url text DEFAULT '/static/avatars/default.png',
  big_photo_url text DEFAULT '/static/avatars/default.png',
  birthdate date,
  biography text DEFAULT 'Биография по данному актеру не заполнена'
);

CREATE INDEX idx_actors_country_id ON public.actors(country_id);
CREATE INDEX idx_actors_name_pattern ON public.actors (first_name text_pattern_ops, second_name text_pattern_ops);

CREATE TABLE public.movie_actors(
  movie_id bigint REFERENCES public.movies(id) ON DELETE CASCADE,
  actor_id bigint REFERENCES public.actors(id) ON DELETE CASCADE,
  PRIMARY KEY (movie_id, actor_id)
);

CREATE INDEX idx_movie_actors_movie_id ON public.movie_actors(movie_id);
CREATE INDEX idx_movie_actors_actor_id ON public.movie_actors(actor_id);

CREATE TABLE public.movie_genres(
  movie_id bigint REFERENCES public.movies(id) ON DELETE CASCADE,
  genre_id bigint REFERENCES public.genres(id) ON DELETE CASCADE,
  PRIMARY KEY (movie_id, genre_id)
);

CREATE INDEX idx_movie_genres_movie_id ON public.movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON public.movie_genres(genre_id);

CREATE TABLE public.collections(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text NOT NULL UNIQUE,
  card_url text DEFAULT '/static/collections/default_card.png',
  album_url text DEFAULT '/static/collections/default_poster.png',
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_collections_title ON public.collections(title);
CREATE INDEX idx_collections_created_at ON public.collections(created_at);

CREATE TABLE public.movie_collections(
  movie_id bigint REFERENCES public.movies(id) ON DELETE CASCADE,
  collection_id bigint REFERENCES public.collections(id) ON DELETE CASCADE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (movie_id, collection_id)
);

CREATE INDEX idx_movie_collections_movie_id ON public.movie_collections(movie_id);
CREATE INDEX idx_movie_collections_collection_id ON public.movie_collections(collection_id);

CREATE TABLE public.ratings (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    movie_id bigint NOT NULL REFERENCES public.movies(id) ON DELETE CASCADE,
    episode_id bigint REFERENCES public.episodes(id) ON DELETE CASCADE,
    rating decimal(10, 2),
    created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_id, movie_id, episode_id)
);

CREATE TABLE episodes (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    movie_id bigint NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    season_number int NOT NULL,
    episode_number int NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    release_date date NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (movie_id, season_number, episode_number)
);
