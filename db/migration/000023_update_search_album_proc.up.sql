DROP PROCEDURE `search_albums`;
CREATE PROCEDURE `search_albums`(IN size INT, IN page INT, IN keyword VARCHAR(50))
BEGIN
	SELECT a.*, get_artists_of_album(a.id) as artist 
	FROM albums a
	WHERE a.name LIKE CONCAT('%', keyword, '%')
	LIMIT size
	OFFSET page;
END
