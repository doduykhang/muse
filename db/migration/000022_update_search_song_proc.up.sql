DROP PROCEDURE `search_songs`;
CREATE PROCEDURE `search_songs` (IN size INT, IN page INT, IN keyword VARCHAR(50))
BEGIN
	SELECT s.*, get_artists_of_song(s.id) as artists, get_albums_of_song(s.id) as albums
	FROM songs s
	WHERE s.title LIKE CONCAT('%', keyword, '%')
	LIMIT size
	OFFSET page;
END
