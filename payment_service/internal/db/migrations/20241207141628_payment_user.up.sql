CREATE ROLE schema_editor_role;

GRANT USAGE ON SCHEMA public TO schema_editor_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO schema_editor_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, UPDATE, CREATE ON TABLES TO schema_editor_role;

CREATE USER payment_user WITH PASSWORD 'payment_user_password';

GRANT schema_editor_role TO payment_user;
