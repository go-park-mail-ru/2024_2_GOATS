CREATE TABLE public.actors(
      id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      first_name text NOT NULL,
      second_name text NOT NULL,
      country_id int REFERENCES public.countries(id),
      small_photo_url text DEFAULT '',
      big_photo_url text DEFAULT '',
      birthdate date,
      biography text DEFAULT ''
);

CREATE INDEX idx_actors_country_id ON public.actors(country_id);
