package dao

// 用户表
var SentenceCreateUser = `
CREATE TABLE IF NOT EXISTS user(
	id 				INT PRIMARY KEY AUTO_INCREMENT,
	password_crypto TINYBLOB NOT NULL COMMENT '密码的md5摘要',
	username		VARCHAR(32) NOT NULL UNIQUE,
	introduction	VARCHAR(256) DEFAULT NULL,
	telephone		VARCHAR(32) DEFAULT NULL,
	gender			TINYINT DEFAULT NULL COMMENT '0代表男, 1代表女',
	email			TEXT DEFAULT NULL,
	birthday		VARCHAR(32) DEFAULT NULL,
	avatar_link		TEXT DEFAULT NULL,
	created_at		DATETIME NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 书籍表
var SentenceCreateBook = `
CREATE TABLE IF NOT EXISTS book(
	id 				INT PRIMARY KEY AUTO_INCREMENT,
	name			VARCHAR(32) NOT NULL,
	author			VARCHAR(32) NOT NULL,
	comment_num		INT NOT NULL,
	score			DOUBLE NOT NULL,
	publish_time	DATETIME NOT NULL,
	cover_link		TEXT NOT NULL,
	content_link	TEXT NOT NULL,
	label_string 	TEXT NOT NULL,
	praise_cnt		INT NOT NULL
)  DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 评论表, 包含书评和二级评论
var SentenceCreateComment = `
CREATE TABLE IF NOT EXISTS comment(
	id 				INT PRIMARY KEY AUTO_INCREMENT,
	book_id			INT NOT NULL,
	user_id			INT NOT NULL,
	username		VARCHAR(32) NOT NULL,
	avatar_link		TEXT NULL,
	praise_cnt		INT NOT NULL,
	content			VARCHAR(512) NOT NULL,
	parent			INT NULL,
	annoymous		TINYINT NOT NULL,
	publish_time	DATETIME NOT NULL,
	created_at		DATETIME NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 关系表, 用户关注书籍
var SentenceCreateRelaFocusUserAndBook = `
CREATE TABLE IF NOT EXISTS rela_focus_user_and_book(
	id				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	user_id			INT NOT NULL,
	book_id			INT NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 关系表, 用户点赞书籍
var SentenceCreateRelaPraiseUserAndBook = `
CREATE TABLE IF NOT EXISTS rela_praise_user_and_book(
	id 				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	user_id			INT NOT NULL,
	book_id			INT NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 关系表, 用户点赞评论
var SentenceCreateRelaPraiseUserAndComment = `
CREATE TABLE IF NOT EXISTS rela_praise_user_and_comment(
	id 				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	user_id			INT NOT NULL,
	comment_id		INT NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 标签表
var SentenceCreateLable = `
CREATE TABLE IF NOT EXISTS label (
	id				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name			VARCHAR(32) NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 关系表, 维护每本书的与标签之间的关系
var SentenceCreateRelaBookAndLabel = `
CREATE TABLE IF NOT EXISTS rela_book_and_label (
	id 				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	book_id			INT NOT NULL,
	label_id		INT NOT NULL
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

// 分布式表锁, 一些复杂的操作不得不锁表
var SentenceCreateTableLock = `
CREATE TABLE IF NOT EXISTS table_lock (
	id 				INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
	tablename 		VARCHAR(32) NOT NULL UNIQUE COMMENT 'unique会建立索引, 因此借助tablename查询不会锁表'
) DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin
`

var SentenceDropUser = `DROP TABLE IF EXISTS user`
var SentenceDropBook = `DROP TABLE IF EXISTS book`
var SentenceDropComment = `DROP TABLE IF EXISTS comment`
var SentenceDropRelaFocusUserAndBook = `DROP TABLE IF EXISTS rela_focus_user_and_book`
var SentenceDropRelaPraiseUserAndBook = `DROP TABLE IF EXISTS rela_praise_user_and_book`
var SentenceDropRelaPraiseUserAndComment = `DROP TABLE IF EXISTS rela_praise_user_and_comment`
var SentenceDropLabel = `DROP TABLE IF EXISTS label`
var SentenceDropRelaBookAndLabel = `DROP TABLE IF EXISTS rela_book_and_label`
var SentenceDropTableLock = `DROP TABLE IF EXISTS table_lock`
