CREATE PROCEDURE `search_artists` (IN size INT, IN page INT, IN keyword VARCHAR(50))
BEGIN
	SELECT a.*
	FROM artists a
	WHERE a.name LIKE CONCAT('%', keyword, '%')
	LIMIT size
	OFFSET page;
END
