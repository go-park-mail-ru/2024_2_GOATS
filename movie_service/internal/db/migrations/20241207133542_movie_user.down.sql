DO $$ BEGIN
   IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'schema_reader_role') THEN
      DROP ROLE schema_reader_role;
   END IF;
END $$;

DO $$ BEGIN
   IF EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'movie_user') THEN
      DROP USER movie_user;
   END IF;
END $$;
