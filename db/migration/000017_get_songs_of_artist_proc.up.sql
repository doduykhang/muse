CREATE PROCEDURE `get_songs_of_artist` (IN artist_id INT)
BEGIN
	SELECT s.*, get_albums_of_song(s.id) as albums
	FROM songs s
	WHERE s.id IN (
		SELECT sa.song_id
		FROM songs_artists sa
		WHERE sa.artist_id = artist_id
	);
END
