CREATE TABLE shortLinks (
	id SERIAL PRIMARY KEY,
	actual_link VARCHAR(255) NOT NULL,
	short_link VARCHAR(255) NOT NULL,
    number_of_uses int DEFAULT(0),
    last_use TEXT
);

CREATE OR REPLACE FUNCTION getDate() RETURNS TEXT AS 
$BODY$
BEGIN
RETURN CURRENT_TIMESTAMP;
END
$BODY$
LANGUAGE 'plpgsql' ;

CREATE OR REPLACE FUNCTION addShortLink(
    actual VARCHAR(255),
    short VARCHAR(255)
    ) RETURNS void AS
$BODY$
BEGIN
    INSERT INTO shortLinks (actual_link, short_link, last_use) 
        VALUES(actual, short, getDate());
    COMMIT;
END
$BODY$
LANGUAGE 'plpgsql' ;