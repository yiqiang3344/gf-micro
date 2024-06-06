// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IMq interface {
		IsTooManyRequest(err error) bool
		// SendMsgs 发送常用类型消息，包括Normal、FIFO、Delay。
		SendMsgs(ctx context.Context, topicType TopicType, topic string, messages []string, sendOptionFunc ...SendOptionFunc) (err error)
	}
)

var (
	localMq IMq
)

func Mq() IMq {
	if localMq == nil {
		panic("implement not found for interface IMq, forgot register?")
	}
	return localMq
}

func RegisterMq(i IMq) {
	localMq = i
}
