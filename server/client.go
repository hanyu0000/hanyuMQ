package server

import "hanyuMQ/kitex_gen/api/client_operations"

type Group struct {
	topics    []*Topic //表示一个 指向 Topic 类型的指针数组
	consumers map[string]*client_operations.Client
}
