CREATE PROCEDURE `get_songs_of_albums` (IN album_id INT)
BEGIN
	SELECT s.*, get_artists_of_song(s.id) as artists
	FROM songs s
	WHERE s.id IN (
		SELECT sa.song_id
		FROM songs_albums sa
		WHERE sa.album_id = album_id
	);
END
