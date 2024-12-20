INSERT INTO public.seasons (movie_id, season_number, title, description, release_date)
VALUES
  (1, 1, 'Игра в Кальмара 1', 'Бойня за деньги', '2023-09-25'),
  (2, 1, 'Бумажный дом 1', 'Бумажная магия', '2024-01-25');

INSERT INTO public.episodes (season_id, episode_number, title, description, release_date, rating, preview_url, video_url)
VALUES
  (3, 1, 'Игра в Кальмара 1 сезон 1 серия', '', '2024-09-25', 9.5, '/static/movies_all/squid-game/season_1/episode-1/preview1.webp', '/static/movies_all/squid-game/season_1/episode-1/squid-game-1.mp4'),
  (3, 2, 'Игра в Кальмара 1 сезон 2 серия', '', '2024-09-25', 7.9, '/static/movies_all/squid-game/season_1/episode-2/preview2.webp', '/static/movies_all/squid-game/season_1/episode-2/squid-game-2.mp4'),
  (3, 3, 'Игра в Кальмара 1 сезон 3 серия', '', '2024-09-25', 8.1, '/static/movies_all/squid-game/season_1/episode-3/preview3.webp', '/static/movies_all/squid-game/season_1/episode-3/squid-game-3.mp4'),
  (3, 4, 'Игра в Кальмара 1 сезон 4 серия', '', '2024-09-25', 8.3, '/static/movies_all/squid-game/season_1/episode-4/preview4.webp', '/static/movies_all/squid-game/season_1/episode-4/squid-game-4.mp4'),
  (4, 1, 'Бумажный дом 2 сезон 1 серия', '', '2024-09-25', 9.0, '/static/movies_all/paper_house/season-1/episode-1/paper_house-1_preview.webp', '/static/movies_all/paper_house/season-1/episode-1/paper_house-1.mp4'),
  (4, 2, 'Бумажный дом 2 сезон 2 серия', '', '2024-01-25', 9.4, '/static/movies_all/paper_house/season-1/episode-2/paper_house-2_preview.webp', '/static/movies_all/paper_house/season-1/episode-2/paper_house-2.mp4'),
  (4, 3, 'Бумажный дом 2 сезон 3 серия', '', '2024-01-25', 8.7, '/static/movies_all/paper_house/season-1/episode-3/paper_house-3_preview.webp', '/static/movies_all/paper_house/season-1/episode-3/paper_house-3.mp4'),
  (4, 4, 'Бумажный дом 2 сезон 4 серия', '', '2024-01-25', 8.6, '/static/movies_all/paper_house/season-1/episode-4/paper_house-4_preview.webp', '/static/movies_all/paper_house/season-1/episode-4/paper_house-4.mp4'),
  (4, 5, 'Бумажный дом 2 сезон 5 серия', '', '2024-01-25', 7.8, '/static/movies_all/paper_house/season-1/episode-5/paper_house-5_preview.webp', '/static/movies_all/paper_house/season-1/episode-5/paper_house-5.mp4');


UPDATE movies SET short_description = 'Захватывающий сериал о преступной группировке, совершающей дерзкий захват Королевского монетного двора Испании'
WHERE movies.id = 2;

DELETE FROM movie_actors
WHERE id = 69;

INSERT INTO movie_actors (movie_id, actor_id)
VALUES
  (24, 35);
