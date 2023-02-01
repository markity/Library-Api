package errorcodes

// 发布书评时, 没有找到这本书
var ErrorCommentBookNotFoundCode = 10600
var ErrorCommentBookNotFoundMsg = "no such a book to be commented"

// 发布二级评论时, 没有找到这个评论
var ErrorCommentCommentNotFoundCode = 10700
var ErrorCommentCommentFoundMsg = "no such a book to be commented"

var ErrorNoSuchCommentToUpdateCode = 10800
var ErrorNoSuchCommentToUpdateMsg = "no such a comment to be updated"
var ErrorNoPermissionToUpdateCommentCode = 10801
var ErrorNoPermissionToUpdateCommentMsg = "no permission to update the comment"

var ErrorNoSuchCommentToDeleteCode = 10900
var ErrorNoSuchCommentToDeleteMsg = "no such a comment to be deleted"
var ErrorNoPermissionToDeleteCommentCode = 10901
var ErrorNoPermissionToDeleteCommentMsg = "no permission to delete the comment"

var ErrorNoSuchBookToListCommentsCode = 11000
var ErrorNoSuchBookToListCommentsMsg = "no such a book to list comments"

var ErrorAlreadyPraisedCommentCode = 11200
var ErrorAlreadyPraisedCommentMsg = "you already praised the comment"
