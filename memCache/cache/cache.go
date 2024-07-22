package cache

import "time"

type Cache interface {
	//size :1kb 100kb 1mb 2mb 1gb
	SetMaxMemory(size string) bool

	//将value写入缓存
	Set(key string, val interface{}, expire time.Duration) bool

	//根据key获取value
	Get(key string) (interface{}, bool)

	//删除key值
	Del(key string) bool

	//判断key是否存在
	Exists(key string) bool

	//清空所有key
	Flush() bool

	//获取缓存中所有key的数量
	Keys() int64
}
