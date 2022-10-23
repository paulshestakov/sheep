CREATE TABLE IF NOT EXISTS service_meta (
    meta_key varchar(256) NOT NULL,
    meta_value jsonb,
    CONSTRAINT service_meta_pkey PRIMARY KEY (meta_key)
);

INSERT INTO service_meta (meta_key, meta_value)
SELECT 'db-version', '"0.0.0"'
WHERE NOT EXISTS (
    SELECT meta_key FROM service_meta WHERE meta_key = 'db-version'
);
