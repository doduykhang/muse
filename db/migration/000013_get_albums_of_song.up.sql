CREATE FUNCTION get_albums_of_song (id int)
RETURNS TEXT DETERMINISTIC
BEGIN
	DECLARE result TEXT;
	SET result = '';
    
	SELECT coalesce(json_arrayagg(json_object("id", a.id, "name", a.name)), "[]") as albums into result
	FROM albums a
	WHERE a.id IN (
		SELECT aa.album_id
		FROM songs_albums aa 
		WHERE aa.song_id = id
	);
    
    	RETURN result;
END 
