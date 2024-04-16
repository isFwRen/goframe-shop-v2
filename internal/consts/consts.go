package consts

const (
	ContextKey        = "ContextKey" // 上下文变量存储键名，前后端系统共享
	CacheModeRedis    = 1            // 缓存模式 1 gcache 2 gredis 3 fileCache
	BackendServerName = "开源电商系统"
	ErrLoginFailMsg   = "登录失败，账号或密码错误"
	GTokenAdminPrefix = "Admin:" //gtoken登录 管理后台 前缀区分
	MultiLogin        = true
	//for 登录相关
	TokenType      = "Bearer"
	GTokenExpireIn = 10 * 24 * 60 * 60
	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	CtxAdminRoleIds = "CtxAdminRoleIds"
)
