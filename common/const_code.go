package common

const (
	SUCCESS        = 200
	SERVER_ERROR          = 500
	INVALID_TOKEN = 302
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)