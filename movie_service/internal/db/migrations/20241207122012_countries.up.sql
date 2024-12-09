CREATE TABLE public.countries(
     id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
     title text NOT NULL,
     code text NOT NULL,
     flag_url text DEFAULT ''
);

CREATE INDEX idx_countries_title ON public.countries(title);
CREATE INDEX idx_countries_code ON public.countries(code);
