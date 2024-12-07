DO $$ BEGIN
   IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'schema_editor_role') THEN
      DROP ROLE schema_editor_role;
   END IF;
END $$;

DO $$ BEGIN
   IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'payment_user') THEN
      DROP USER payment_user;
   END IF;
END $$;
