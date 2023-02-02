package model

import (
	"time"
)

type FocusListItem struct {
	BookID      int64     `json:"book_id"`
	Name        string    `json:"name"`
	PublishTime time.Time `json:"publish_time"`
	ContentLink string    `json:"content_link"`
}

type User struct {
	ID int64
	// 0男 1女
	Gender int
	// 密码的md5摘要
	PasswordCrypto string
	Nickname       string
	QQ             string
	Birthday       time.Time
	Email          string
	// 头像的链接
	AvatarLink string
	// 简介
	Introduction string
	Phone        string
}

type Book struct {
	ID          int64     `json:"book_id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	CommentNum  int64     `json:"comment_num"`
	Score       float64   `json:"score"`
	PublishTime time.Time `json:"publish_time"`
	CoverLink   string    `json:"cover_link"`
	ContentLink string    `json:"content_link"`
	Label       string    `json:"label"`
	PraiseCnt   int64     `json:"praise_cnt"`
}

// 细节, 评论分为二级评论和顶级评论
//		顶级评论是评价书籍的, Parent=NULL
// 		二级评论是在别人评价上追评的, Parent!=NULL

type CommentWithParise struct {
	Comment Comment `json:"comment"`
	Praise  *bool   `json:"praise,omitempty"`
}

type Comment struct {
	ID           int64                `json:"id"`
	BookID       int64                `json:"book_id"`
	Username     *string              `json:"username,omitempty"`
	AvatarLink   *string              `json:"avatar_link"`
	PraiseCnt    int64                `json:"praise_cnt"`
	Content      string               `json:"content"`
	SenderUserID int64                `json:"-"`
	Parent       *int64               `json:"-"`
	Anonymous    bool                 `json:"annoymous"`
	PublishTime  time.Time            `json:"publish_time"`
	SonComments  []*CommentWithParise `json:"son"`
}

// 这张表维护用户和书的一对多关系, 用来说明
// 一个用户是否关注了一本书
type RelationFocusUserAndBook struct {
	UserID int64
	BookID int64
}

// 这张表维护用户和书的一对多关系, 用来说明
// 一个用户是否点赞了一本书
type RelationPraiseUserAndBook struct {
	UserID int64
	BookID int64
}
