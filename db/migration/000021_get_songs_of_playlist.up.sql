CREATE PROCEDURE `get_songs_of_playlist` (IN playlist_id INT)
BEGIN
	SELECT s.*, get_artists_of_song(s.id) as artists, get_albums_of_song(s.id) as albums
	FROM songs s
	WHERE s.id IN (
		SELECT sp.song_id
		FROM songs_playlists sp
		WHERE sp.playlist_id = playlist_id
	);
END
