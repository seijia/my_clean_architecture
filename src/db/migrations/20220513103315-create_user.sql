
-- +migrate Up
CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name varchar(255) DEFAULT NULL COMMENT 'ユーザ名',
   age varchar(255) DEFAULT NULL COMMENT '年齢',
   created_at datetime DEFAULT NULL COMMENT '作成日時',
   updated_at datetime DEFAULT NULL COMMENT '更新日時',
   deleted_at timestamp NULL DEFAULT NULL COMMENT '削除日時',
   INDEX user_id (id),
   PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';


insert into users(name, age) values ("user1", 28);
insert into users(name, age) values ("user2", 29);
-- +migrate Down
 DROP TABLE IF EXISTS users;