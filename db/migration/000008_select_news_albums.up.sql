CREATE PROCEDURE `select_new_albums` (IN size INT, IN page INT)
BEGIN
	SELECT albums.*, (
		SELECT coalesce(json_arrayagg(json_object("id", a.id, "name", a.name)), "[]")
		FROM artists a
		WHERE a.id IN (
			SELECT aa.artist_id
			FROM artists_albums aa 
			WHERE aa.album_id = albums.id
		)
	) AS artists,
	(
		SELECT album_types.name 
		FROM album_types
		WHERE album_types.id = albums.album_type_id
	) AS album_type
	FROM albums
	ORDER BY albums.created_at DESC
	LIMIT size
	OFFSET page;
END
