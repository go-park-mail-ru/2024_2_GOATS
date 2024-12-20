INSERT INTO public.movie_genres (movie_id, genre_id)
VALUES
  (11, 3),
  (7, 3),
  (8, 3),
  (26, 3),
  (13, 3),
  (15, 3),

  (14, 4),
  (35, 4),

  (34, 14),
  (35, 14),
  (36, 14),
  (31, 14),
  (30, 14),

  (34, 13),
  (37, 13),
  (19, 13),

  (37, 11),
  (14, 11);

DELETE FROM movie_genres
WHERE movie_id = 34 and genre_id = 16 or movie_id = 3 and genre_id = 6;
