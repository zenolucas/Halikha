DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id         INT AUTO_INCREMENT NOT NULL,
  username      VARCHAR(128) NOT NULL,
  usertype     VARCHAR(8) NOT NULL,
  password      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users
  (username, usertype, password)
VALUES
  ('user1', 'artist', 'pass'),
  ('user2', 'customer', 'pass');
