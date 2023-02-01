package service

import (
	"library-api/dao"
	"library-api/model"
	errorcodes "library-api/util/error_codes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt鉴权错误
func RespJWTError(ctx *gin.Context) {
	respStruct := errorcodes.BasicErrorResp{
		Status: errorcodes.ErrorInvalidUserTokenCode,
		Info:   errorcodes.ErrorInvalidUserTokenMsg,
	}

	ctx.JSON(http.StatusOK, &respStruct)
}

// 服务暂时不可用
func RespServiceNotAvailabelError(ctx *gin.Context) {
	respStruct := errorcodes.BasicErrorResp{
		Status: errorcodes.ErrorServiceNotAvailabelCode,
		Info:   errorcodes.ErrorServiceNotAvailabelMsg,
	}

	ctx.JSON(http.StatusOK, &respStruct)
}

// 入参错误
func RespInvalidParaError(ctx *gin.Context) {
	respStruct := errorcodes.BasicErrorResp{
		Status: errorcodes.ErrorInvalidInputParametersCode,
		Info:   errorcodes.ErrorInvalidInputParametersMsg,
	}

	ctx.JSON(http.StatusOK, &respStruct)
}

func MustResetTables() {
	_, err := dao.DB.Exec(dao.SentenceDropBook)
	if err != nil {
		log.Panicf("failed to drop table book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropComment)
	if err != nil {
		log.Panicf("failed to drop table commment: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropRelaFocusUserAndBook)
	if err != nil {
		log.Panicf("failed to drop table rela_focus_user_and_book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropRelaPraiseUserAndBook)
	if err != nil {
		log.Panicf("failed to drop table rela_praise_user_and_book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropRelaPraiseUserAndComment)
	if err != nil {
		log.Panicf("failed to drop table rela_praise_user_and_comment: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropUser)
	if err != nil {
		log.Panicf("failed to drop table user: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropLabel)
	if err != nil {
		log.Panicf("failed to drop table label: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropRelaBookAndLabel)
	if err != nil {
		log.Panicf("failed to drop table rela_book_and_label: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceDropTableLock)
	if err != nil {
		log.Panicf("failed to drop table table_lock: %v\n", err)
	}

	// --------------建表--------------
	_, err = dao.DB.Exec(dao.SentenceCreateBook)
	if err != nil {
		log.Panicf("failed to create table book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateComment)
	if err != nil {
		log.Panicf("failed to create table comment: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateRelaFocusUserAndBook)
	if err != nil {
		log.Panicf("failed to create table rela_focus_user_and_book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateRelaPraiseUserAndBook)
	if err != nil {
		log.Panicf("failed to create table rela_praise_user_and_book: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateRelaPraiseUserAndComment)
	if err != nil {
		log.Panicf("failed to create table rela_praise_user_and_comment: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateUser)
	if err != nil {
		log.Panicf("failed to create table user: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateLable)
	if err != nil {
		log.Panicf("failed to create table label: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateRelaBookAndLabel)
	if err != nil {
		log.Panicf("failed to create table rela_book_and_label: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateTableLock)
	if err != nil {
		log.Panicf("failed to create table table_lock: %v\n", err)
	}

	// -------------准备数据------------------
	// 书籍数据是预先准备的, 可能是来自爬虫爬取的, 此处预先准备一些
	b := model.Book{}
	b.Name = "傲慢与偏见"
	b.Author = "in id mollit labore"
	b.CommentNum = 0
	b.ContentLink = "http://127.0.0.1/test1"
	b.CoverLink = "http://127.0.0.1/picture_url"
	b.Label = "悬疑,恐怖,暴力"
	b.PublishTime = time.Now()
	b.Score = 7.3
	b.PraiseCnt = 0
	_, err = dao.DB.Exec(`INSERT INTO book(name,author,comment_num,content_link,
		 cover_link,label_string,publish_time,score,praise_cnt) VALUES(?,?,?,?,?,?,?,?,?)`,
		b.Name, b.Author, b.CommentNum, b.ContentLink, b.CoverLink, b.Label, b.PublishTime, b.Score, b.PraiseCnt)
	if err != nil {
		log.Panicf("failed to insert data: %v\n", err)
	}
	// 创建label数据
	_, err = dao.DB.Exec("INSERT INTO label(name) VALUES(?), (?), (?)", "悬疑", "恐怖", "暴力")
	if err != nil {
		log.Fatalf("failed to prepare data: %v\n", err)
	}
	// 将label与书籍关联, 这是平时就需要维护的一张表, 便于借助label
	_, err = dao.DB.Exec("INSERT INTO rela_book_and_label(book_id, label_id) VALUES(1, 1), (1, 2), (1, 3)")
	if err != nil {
		log.Fatalf("failed to prepare data: %v\n", err)
	}
	// 准备分布式表锁
	_, err = dao.DB.Exec("INSERT INTO table_lock(tablename) VALUES(?), (?), (?)", "rela_praise_user_and_book", "rela_focus_user_and_book", "rela_praise_user_and_comment")
	if err != nil {
		log.Fatalf("failed to prepare data: %v\n", err)
	}

	log.Println("reset database ok")
}
