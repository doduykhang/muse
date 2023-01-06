DROP PROCEDURE `select_my_playlist`;
CREATE PROCEDURE `select_my_playlist`(IN size INT,IN page INT, IN user_id INT)
BEGIN
	SELECT * 
	from playlists 
	where user_id = user_id
	LIMIT size
	OFFSET page;
END
