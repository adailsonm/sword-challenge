CREATE TABLE tasks (
	id			int				NOT NULL	AUTO_INCREMENT,
  user_id     int         NOT NULL,
	summary		  text(2500)  NOT NULL,
  created_at	timestamp		NOT NULL,
	updated_at	timestamp		NOT NULL,
	PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);