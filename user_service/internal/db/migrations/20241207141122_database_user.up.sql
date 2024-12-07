CREATE ROLE schema_editor_role;

GRANT USAGE ON SCHEMA public TO schema_editor_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO schema_editor_role;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, UPDATE, CREATE, DELETE ON TABLES TO schema_editor_role;

CREATE USER user_service_user WITH PASSWORD 'user_service_user_password';

GRANT schema_editor_role TO user_service_user;
