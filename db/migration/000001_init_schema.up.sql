CREATE TABLE IF NOT EXISTS accounts (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	active TINYINT(1) NOT NULL DEFAULT 0,
	banned TINYINT(1) NOT NULL DEFAULT 0,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	gender TINYINT(1) NOT NULL DEFAULT 0,
	birthdate DATE NOT NULL DEFAULT 0,
	account_id BIGINT NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS playlists (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	name VARCHAR(255) NOT NULL,
	image VARCHAR(255) NOT NULL,
	user_id BIGINT NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS songs (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	title VARCHAR(255) NOT NULL,
	duration VARCHAR(255) NOT NULL,
	release_date DATE NOT NULL,
	image VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS songs_playlists (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	song_id BIGINT NOT NULL,
	playlist_id BIGINT NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (song_id) REFERENCES songs(id),
	FOREIGN KEY (playlist_id) REFERENCES playlists(id)
);

CREATE TABLE IF NOT EXISTS artists (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	image VARCHAR(255) NOT NULL,
	birthdate DATE NOT NULL,
	gender TINYINT(1) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS songs_artists (
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	song_id BIGINT NOT NULL,
	artist_id BIGINT NOT NULL,
	PRIMARY KEY (song_id, artist_id),
	FOREIGN KEY (song_id) REFERENCES songs(id),
	FOREIGN KEY (artist_id) REFERENCES artists(id)
);

CREATE TABLE IF NOT EXISTS album_types (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS albums (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	name VARCHAR(255) NOT NULL,
	image VARCHAR(255) NOT NULL,
	release_date DATE NOT NULL,
	album_type_id BIGINT NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (album_type_id) REFERENCES album_types(id)
);

CREATE TABLE IF NOT EXISTS songs_albums (
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	song_id BIGINT NOT NULL,
	album_id BIGINT NOT NULL,
	PRIMARY KEY (song_id, album_id),
	FOREIGN KEY (song_id) REFERENCES songs(id),
	FOREIGN KEY (album_id) REFERENCES albums(id)
);

CREATE TABLE IF NOT EXISTS genres (
	id BIGINT NOT NULL AUTO_INCREMENT,
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS songs_genres (
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	deleted_at DATETIME NOT NULL,
	song_id BIGINT NOT NULL,
	genre_id BIGINT NOT NULL,
	PRIMARY KEY (song_id, genre_id),
	FOREIGN KEY (song_id) REFERENCES songs(id),
	FOREIGN KEY (genre_id) REFERENCES genres(id)
);

