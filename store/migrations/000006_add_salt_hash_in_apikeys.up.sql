ALTER TABLE api_keys
ADD COLUMN salt_hash varchar(255) AFTER api_key