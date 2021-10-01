CREATE TABLE anime (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) UNIQUE NOT NULL,
    description TEXT,
    num_episodes INTEGER,
    img VARCHAR(255)
);

CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(500) NOT NULL
);

CREATE TABLE user_account_anime (
    id SERIAL PRIMARY KEY,
    user_account_id INTEGER NOT NULL REFERENCES user_account(id),
    anime_id INTEGER NOT NULL REFERENCES anime(id),
    status VARCHAR(25),
    last_episode_seen INTEGER
);
