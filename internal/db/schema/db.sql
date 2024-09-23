DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS collections CASCADE;
DROP TABLE IF EXISTS movie_collections CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS movie_genres CASCADE;
DROP TABLE IF EXISTS countries CASCADE;

CREATE TABLE public.genres(
  id serial PRIMARY KEY,
  title varchar NOT NULL UNIQUE,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT valid_title CHECK ( LENGTH(title) >= 3 )
);

CREATE INDEX idx_genres_title ON public.genres(title);

INSERT INTO public.genres (title)
VALUES
  ('Ужасы'),
  ('Комедия'),
  ('Психология'),
  ('Фэнтези'),
  ('Триллер'),
  ('Романтика');

CREATE TABLE public.users(
  id serial PRIMARY KEY,
  username varchar UNIQUE,
  email varchar NOT NULL UNIQUE,
  password_hash varchar NOT NULL,
  sex varchar CHECK(sex IN ('male', 'female', 'other', 'secret')),
  birthdate date,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT valid_username CHECK ( LENGTH(username) >= 3 and LENGTH(username) <= 20 ),
  CONSTRAINT valid_email CHECK ( email ~* '\S*@\S*')
);

CREATE INDEX idx_users_email ON public.users(email);
CREATE INDEX idx_users_birthdate ON public.users(birthdate);
CREATE INDEX idx_users_created_at ON public.users(created_at);
CREATE INDEX idx_users_updated_at ON public.users(updated_at);

CREATE TABLE public.countries(
  id serial PRIMARY KEY,
  title varchar NOT NULL,
  code varchar NOT NULL,
  flag_url varchar DEFAULT ''
);

CREATE INDEX idx_countries_title ON public.countries(title);
CREATE INDEX idx_countries_code ON public.countries(code);

INSERT INTO public.countries (title, code)
VALUES
  ('Россия', 'rus'),
  ('США', 'usa'),
  ('Великобритания', 'uk');

CREATE TABLE public.movies(
  id serial PRIMARY KEY,
  title varchar NOT NULL,
  description text NOT NULL,
  card_url varchar DEFAULT '',
  album_url varchar DEFAULT '',
  release_date date NOT NULL,
  rating decimal(10,2) DEFAULT '0.0',
  movie_type varchar CHECK(movie_type IN ('film', 'serial')),
  country_id int REFERENCES public.countries(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_movies_title ON public.movies(title);
CREATE INDEX idx_movies_movie_type ON public.movies(movie_type);
CREATE INDEX idx_movies_release_date ON public.movies(release_date);
CREATE INDEX idx_movies_country_id ON public.movies(country_id);

INSERT INTO public.movies (title, country_id, description, release_date, movie_type)
VALUES
  ('Оно', 2, 'Когда в городке Дерри, штат Мэн, начинают пропадать дети, несколько ребят сталкиваются со своими величайшими страхами и вынуждены помериться силами со злобным клоуном Пеннивайзом, чьи проявления жестокости и список жертв уходят в глубь веков.', '2017-09-07', 'film'),
  ('Очень странные дела: 1 сезон', 2, '«О́чень стра́нные дела́» — американский научно-фантастический драматический сериал с элементами ужасов, созданный братьями Даффер для стриминг-сервиса Netflix. Дафферы выступают шоураннерами, а также исполнительными продюсерами сериала вместе с Шоном Леви и Дэном Коэном.', '2016-07-15', 'serial'),
  ('Интерстеллар', 2, 'Наше время на Земле подошло к концу, команда исследователей берет на себя самую важную миссию в истории человечества; путешествуя за пределами нашей галактики, чтобы узнать есть ли у человечества будущее среди звезд.', '2014-11-06', 'film'),
  ('Трансформеры', 2,'В течение многих столетий две расы роботов-инопланетян — Автоботы и Десептиконы — вели войну, ставкой в которой была судьба Вселенной. И вот война докатилась до Земли. В то время, когда силы зла ищут ключ к верховной власти, наш последний шанс на спасение находится в руках юного землянина. Единственное, что стоит между несущими зло Десептиконами и высшей властью — это ключ, находящийся в руках простого парнишки.', '2007-07-04', 'film'),
  ('Барби', 3, 'Барби выгоняют из Барбиленда, потому что она не соответствует его нормам красоты. Тогда она начинает новую жизнь в реальном мире, где обнаруживает, что совершенства можно достичь только благодаря внутренней гармонии.', '2023-07-20', 'film');

CREATE TABLE public.movie_genres(
  id serial PRIMARY KEY,
  movie_id int REFERENCES public.movies(id),
  genre_id int REFERENCES public.genres(id)
);

CREATE INDEX idx_movie_genres_movie_id ON public.movie_genres(movie_id);
CREATE INDEX idx_movie_genres_genre_id ON public.movie_genres(genre_id);

INSERT INTO public.movie_genres (movie_id, genre_id)
VALUES
  (1, 1),
  (1, 3),
  (1, 5),
  (2, 1),
  (2, 4),
  (3, 4),
  (3, 5),
  (4, 4),
  (5, 6),
  (5, 2),
  (5, 3);

CREATE TABLE public.collections(
  id serial PRIMARY KEY,
  title varchar NOT NULL UNIQUE,
  card_url varchar,
  album_url varchar,
  conditions json,
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT valid_title CHECK ( LENGTH(title) >= 3 )
);

INSERT INTO public.collections (title, card_url)
VALUES
  ('Классика 10-х годов', ''),
  ('Женская правда', ''),
  ('ОООЧЕНЬ страшно', '');

CREATE INDEX idx_collections_title ON public.collections(title);
CREATE INDEX idx_collections_created_at ON public.collections(created_at);

CREATE TABLE public.movie_collections(
  id serial PRIMARY KEY,
  movie_id int REFERENCES public.movies(id),
  collection_id int REFERENCES public.collections(id),
  created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO public.movie_collections (movie_id, collection_id)
VALUES
  (1, 3),
  (2, 3),
  (3, 1),
  (4, 1),
  (2, 2),
  (5, 2);

CREATE INDEX idx_movie_collections_movie_id ON public.movie_collections(movie_id);
CREATE INDEX idx_movie_collections_collection_id ON public.movie_collections(collection_id);
