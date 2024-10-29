CREATE TABLE public.actors(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name TEXT NOT NULL,
  second_name TEXT NOT NULL,
  country_id BIGINT NOT NULL REFERENCES public.countries(id),
  small_photo_url TEXT DEFAULT '/static/avatars/default.png',
  big_photo_url TEXT DEFAULT '/static/avatars/default.png',
  birthdate DATE,
  biography TEXT DEFAULT 'Биография по данному актеру не заполнена'
);

CREATE INDEX idx_actors_country_id ON public.actors(country_id);
CREATE INDEX idx_actors_name_pattern ON public.actors (first_name text_pattern_ops, second_name text_pattern_ops);
