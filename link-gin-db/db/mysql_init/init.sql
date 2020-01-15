CREATE DATABASE if not exists db_docker_link collate utf8mb4;
use db_docker_link;

create table if not exists tb_user
(
	id int unsigned auto_increment primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	username varchar(100) null,
	password varchar(80) null,
	email varchar(100) null,
	phone varchar(20) null,
	sex tinyint null,
	address varchar(500) null,
	is_available tinyint null,
	last_login datetime null,
	login_ip varchar(20) null,
	constraint uix_tb_user_username unique (username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

create index idx_tb_user_deleted_at on tb_user (deleted_at);