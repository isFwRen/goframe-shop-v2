package consts

const (
	ContextKey           = "ContextKey" // 上下文变量存储键名，前后端系统共享
	CacheModeRedis       = 1            // 缓存模式 1 gcache 2 gredis 3 fileCache
	BackendServerName    = "开源电商系统"
	ErrLoginFailMsg      = "登录失败，账号或密码错误"
	GTokenAdminPrefix    = "Admin:" //gtoken登录 管理后台 前缀区分
	GTokenFrontendPrefix = "User:"  //gtoken登录 前台用户 前缀区分
	MultiLogin           = true
	//for 登录相关
	TokenType      = "Bearer"
	GTokenExpireIn = 10 * 24 * 60 * 60
	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	CtxAdminRoleIds = "CtxAdminRoleIds"
	//文章相关
	ArticleIsAdmin = 1 //管理员发布
	ArticleIsUser  = 2 //用户发布
	//统一管理错误提示
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrSecretAnswerMsg      = "密保问题不正确"
	ResourcePermissionFail  = "没有权限操作"
	FrontendMultiLogin      = true
	//for user
	CtxUserId        = "CtxUserId"
	CtxUserName      = "CtxUserName"
	CtxUserAvatar    = "CtxUserAvatar"
	CtxUserSex       = "CtxUserSex"
	CtxUserSign      = "CtxUserSign"
	CtxUserStatus    = "CtxUserStatus"
	RefundStatusWait = 1
	//收藏相关
	CollectionTypeGoods   = 1
	CollectionTypeArticle = 2
	//收货地址相关
	ProvincePid = 1
)
