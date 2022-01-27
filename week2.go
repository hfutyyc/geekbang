package main

import (
	"database/sql"
	"fmt"
)

func main() {
	content, err := dao()
	if err != nil {
		return
	}
	// 其他业务逻辑
	return
}

func service() (string, error) {
	content, err := dao()
	if err != nil {
		return "", err
	}
	// 其他业务逻辑
	return content, nil
}

// 不需要在dao层把sql.ErrNoRows Warp后抛给上层，直接抛error就行
func dao() (string, error) {
	content, err := sqlFunc()
	if err != nil {
		// 打印错误日志
		fmt.Printf("xxx err: %v", err)
		return "", err
	}
	// 处理数据
	return content, nil
}

func sqlFunc() (string, error) {
	return "", sql.ErrNoRows
}
