DELETE FROM episodes
WHERE episodes.season_id in (3, 4);

DELETE FROM seasons
WHERE seasons.id in (3, 4);
