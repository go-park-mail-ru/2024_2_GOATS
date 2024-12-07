CREATE ROLE schema_reader_role;

GRANT USAGE ON SCHEMA public TO schema_reader_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO schema_reader_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO schema_reader_role;

CREATE USER movie_user WITH PASSWORD 'movie_user_password';

GRANT schema_reader_role TO movie_user;
