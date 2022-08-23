CREATE DATABASE account_service_app_project;

USE account_service_app_project;

CREATE TABLE users (
  no_telp INT NOT NULL PRIMARY KEY,
  password_user VARCHAR(15) NOT NULL,
  name_user VARCHAR(20) NOT NULL,
  gender VARCHAR(10) NOT NULL,
  balance BIGINT NOT NULL,
  currency VARCHAR(10),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE top_up (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  account_telp INT NOT NULL,
  amount BIGINT NOT NULL COMMENT "must be positive",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_entries_user FOREIGN KEY (account_telp) REFERENCES users(no_telp)
);

CREATE TABLE transfers (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  from_account_telp INT NOT NULL,
  to_account_telp INT NOT NULL,
  amount BIGINT NOT NULL COMMENT "must be positive",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_transferfrom_user FOREIGN KEY (from_account_telp) REFERENCES users(no_telp),
  CONSTRAINT fk_transferto_user FOREIGN KEY (to_account_telp) REFERENCES users(no_telp)
);

ALTER TABLE top_up  
DROP FOREIGN KEY  fk_entries_user ;
ALTER TABLE top_up  
ADD CONSTRAINT  fk_entries_user FOREIGN KEY ( account_telp) REFERENCES  users  ( no_telp )
ON DELETE CASCADE
ON UPDATE CASCADE;
  
ALTER TABLE transfers 
DROP FOREIGN KEY fk_transferfrom_user;
ALTER TABLE  transfers  
ADD CONSTRAINT  fk_transferfrom_user FOREIGN KEY ( from_account_telp ) REFERENCES  users  ( no_telp )
ON DELETE CASCADE
ON UPDATE CASCADE;

ALTER TABLE transfers 
DROP FOREIGN KEY fk_transferto_user;
ALTER TABLE  transfers  
ADD CONSTRAINT  fk_transferto_user FOREIGN KEY ( to_account_telp ) REFERENCES  users  ( no_telp )
ON DELETE CASCADE
ON UPDATE CASCADE;




