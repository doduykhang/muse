DROP PROCEDURE `select_my_playlist`;
CREATE PROCEDURE `select_my_playlist`(IN size INT,IN page INT)
BEGIN
	SELECT * 
	from playlists 
	where user_id = 1
	LIMIT size
	OFFSET page;
END
