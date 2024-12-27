CREATE DATABASE greenlight;
\c greenlight
CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';
CREATE EXTENSION IF NOT EXISTS citext;

GRANT USAGE ON SCHEMA public TO greenlight;
GRANT CREATE ON SCHEMA public TO greenlight;