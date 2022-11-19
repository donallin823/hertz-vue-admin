package system

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "github.com/edufriendchen/hertz-vue-admin/server/api/v1"
	"github.com/edufriendchen/hertz-vue-admin/server/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(h *server.Hertz) {
	userRouter := h.Group("user").Use(middleware.OperationRecord)
	userRouterWithoutRecord := h.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)       // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)        // 获取自身信息
		userRouterWithoutRecord.POST("exportUserInfo", baseApi.ExportUserList) // 导出用户信息
	}
}
