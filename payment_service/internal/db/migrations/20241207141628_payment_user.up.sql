DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_roles WHERE rolname = 'schema_editor_role'
    ) THEN
        CREATE ROLE schema_editor_role;
    END IF;
END $$;

GRANT USAGE ON SCHEMA public TO schema_editor_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO schema_editor_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, UPDATE, CREATE ON TABLES TO schema_editor_role;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_roles WHERE rolname = 'payment_user'
    ) THEN
        CREATE USER payment_user WITH PASSWORD 'payment_user_password';
    END IF;
END $$;

GRANT schema_editor_role TO payment_user;
