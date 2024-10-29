CREATE TABLE public.countries(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title TEXT NOT NULL,
  code TEXT NOT NULL,
  flag_url TEXT DEFAULT '/static/flags/default.png'
);

CREATE INDEX idx_countries_title ON public.countries(title);
CREATE INDEX idx_countries_code ON public.countries(code);
