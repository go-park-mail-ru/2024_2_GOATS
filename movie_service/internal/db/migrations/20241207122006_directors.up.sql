CREATE TABLE public.directors(
         id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
         first_name text NOT NULL,
         second_name text NOT NULL
);
