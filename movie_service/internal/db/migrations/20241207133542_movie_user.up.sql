DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_roles WHERE rolname = 'schema_reader_role'
    ) THEN
        CREATE ROLE schema_reader_role;
    END IF;
END $$;

GRANT USAGE ON SCHEMA public TO schema_reader_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO schema_reader_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO schema_reader_role;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_roles WHERE rolname = 'movie_user'
    ) THEN
        CREATE USER movie_user WITH PASSWORD 'movie_user_password';
    END IF;
END $$;

GRANT schema_reader_role TO movie_user;
