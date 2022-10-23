DO $$
BEGIN
    IF (SELECT 1 FROM service_meta WHERE meta_key = 'db-version' AND meta_value = '"0.0.0"') THEN
        CREATE TABLE IF NOT EXISTS users (
            "id" uuid NOT NULL,
            "name" text NOT NULL,
            PRIMARY KEY(id)
        );

        UPDATE service_meta SET meta_value = '"0.0.1"' WHERE meta_key = 'db-version';
    END IF;
END$$;
