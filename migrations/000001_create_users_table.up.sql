
CREATE TABLE users (
	id			int				NOT NULL	AUTO_INCREMENT,
  name    varchar(50)	 NOT NULL,
  role		varchar(20)		NOT NULL,
	email		varchar(150)	NOT NULL,
	password	varchar(200)	NOT NULL,
	PRIMARY KEY (id),
  created_at	timestamp		NOT NULL,
	updated_at	timestamp		NOT NULL
);