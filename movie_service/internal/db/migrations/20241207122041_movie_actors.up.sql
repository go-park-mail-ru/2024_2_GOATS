CREATE TABLE public.movie_actors(
            id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            movie_id int REFERENCES public.movies(id),
            actor_id int REFERENCES public.actors(id)
);

CREATE INDEX idx_movie_actors_movie_id ON public.movie_actors(movie_id);
CREATE INDEX idx_movie_actors_actor_id ON public.movie_actors(actor_id);
