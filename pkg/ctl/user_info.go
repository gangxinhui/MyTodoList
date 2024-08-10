package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

/*
这三个方法的意义在于利用上下文对象传递和获取用户信息，以便在不同函数或服务之间共享用户数据：

GetUserInfo：方便地从上下文中提取用户信息。如果无法提取，则返回错误。这在需要用户信息的地方非常有用，如认证和授权。
NewContext：创建一个包含用户信息的新上下文。通常用于在处理请求时，将用户信息注入上下文，使后续处理可以方便地访问这些信息。
FromContext：从上下文中提取用户信息的实际实现。GetUserInfo 方法调用它来实现具体的提取逻辑。
通过这些方法，可以在请求处理流程中方便地传递和获取用户信息，确保用户信息在不同的服务和函数之间保持一致和安全。
*/
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
