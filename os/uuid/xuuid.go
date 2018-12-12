package uuid

import "github.com/google/uuid"

// 获取文件或目录信息
func UUid() string {
	u := uuid.New()
	return u.String()
}
