DROP PROCEDURE `select_my_playlist`;
CREATE PROCEDURE `select_my_playlist`(IN size INT,IN page INT, IN _user_id INT)
BEGIN
	SELECT * 
	from playlists 
	where user_id = _user_id
	LIMIT size
	OFFSET page;
END
