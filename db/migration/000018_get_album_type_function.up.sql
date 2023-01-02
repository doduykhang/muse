CREATE FUNCTION get_album_type (id int)
RETURNS TEXT DETERMINISTIC
BEGIN
	DECLARE result TEXT;
	SET result = '';
    
	SELECT a.name INTO result
	FROM album_types a
	WHERE a.id = id;    

    	RETURN result;
END 
