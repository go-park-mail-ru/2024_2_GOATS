CREATE TABLE public.qualities (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    quality quality_enum NOT NULL
);
