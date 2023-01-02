CREATE PROCEDURE `get_albums_of_artist` (IN artist_id INT)
BEGIN
	SELECT a.*, get_album_type(a.album_type_id) as album_type
	FROM albums a
	WHERE a.id IN (
		SELECT aa.album_id
		FROM artists_albums aa
		WHERE aa.artist_id = artist_id
	);
END
