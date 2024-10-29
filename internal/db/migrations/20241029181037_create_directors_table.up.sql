CREATE TABLE public.directors(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name TEXT NOT NULL,
  second_name TEXT NOT NULL
);

CREATE INDEX idx_directors_name_pattern ON public.directors (first_name text_pattern_ops, second_name text_pattern_ops);
