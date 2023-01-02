CREATE FUNCTION get_artists_of_song (id int)
RETURNS TEXT DETERMINISTIC
BEGIN
	DECLARE result TEXT;
	SET result = '';
    
	SELECT coalesce(json_arrayagg(json_object("id", a.id, "name", a.name)), "[]") as artist into result
	FROM artists a
	WHERE a.id IN (
		SELECT aa.artist_id
		FROM songs_artists aa 
		WHERE aa.song_id = id
	);
    
    	RETURN result;
END 
