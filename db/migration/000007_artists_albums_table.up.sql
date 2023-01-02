CREATE TABLE IF NOT EXISTS artists_albums (
	created_at DATETIME,
	created_by VARCHAR(255),
	updated_at DATETIME,
	updated_by VARCHAR(255),
	deleted_at DATETIME,
	deleted_by VARCHAR(255),
	album_id BIGINT NOT NULL,
	artist_id BIGINT NOT NULL,
	PRIMARY KEY (album_id, artist_id),
	FOREIGN KEY (album_id) REFERENCES albums(id),
	FOREIGN KEY (artist_id) REFERENCES artists(id)
);
