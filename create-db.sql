SELECT 'CREATE DATABASE animelist'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'animelist')\gexec
