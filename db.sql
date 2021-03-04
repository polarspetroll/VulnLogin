CREATE TABLE vulnlogin(
  username VARCHAR(30) PRIMARY KEY,
  password VARCHAR(20)
);

INSERT INTO vulnlogin VALUES('admin', 'password');
INSERT INTO vulnlogin VALUES('john', 'pwd*&)(*^*&^)');
INSERT INTO vulnlogin VALUES('admin2', 'PaSs1234');
INSERT INTO vulnlogin VALUES('devteam', 'P@SsW0rD');
