CREATE PROCEDURE select_songs (IN size INT,IN page INT)
BEGIN
SELECT s.*, (
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
LIMIT size
OFFSET page;
END
