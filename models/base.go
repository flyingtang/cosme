package models

type BaseModel struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// 初始化表格的sql语句集合
var tables = []string{`
CREATE TABLE IF NOT EXISTS Account(
   	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
   	username VARCHAR(128) not null unique,
	password VARCHAR(128) not null ,
	nickname VARCHAR(128),
	sex TINYINT,
	email VARCHAR(128),
   	created_at bigint  NOT NULL,
   	updated_at bigint  NOT NULL,
	INDEX username_index(username),
   	PRIMARY KEY (id)
);`}