CREATE PROCEDURE `select_songs_of_genre` (IN size INT, IN page INT, IN genre_id INT)
BEGIN
	SELECT s.*,
		(
			SELECT coalesce(json_arrayagg(json_object("id", a.id, "name", a.name)), "[]")
    			FROM artists a
			WHERE a.id in (
				SELECT sa.artist_id
        			FROM songs_artists sa
        			WHERE sa.song_id = s.id
    			)
		) as artists,
		(
			SELECT coalesce(json_arrayagg(json_object("id", a.id, "name", a.name)), "[]")
    			FROM albums a
			WHERE a.id in (
				SELECT sa.album_id
        			FROM songs_albums sa
        			WHERE sa.song_id = s.id
    		)
		) as albums
	FROM songs s
	WHERE s.id IN (
		SELECT sg.song_id
		FROM songs_genres sg
		WHERE sg.genre_id = genre_id
	)
	LIMIT size
	OFFSET page;
END
