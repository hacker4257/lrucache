
# LRU Cache 实现

这是一个用 Go 语言实现的 LRU (Least Recently Used) 缓存数据结构。

## 项目简介

LRU 缓存是一种常用的缓存淘汰算法,当缓存满时会优先删除最近最少使用的数据。本项目提供了 LRU 缓存的基本实现,包含以下特性:

- 支持设置缓存容量上限
- 基于双向链表和哈希表实现 O(1) 的查询和更新操作
- 线程安全的实现
- 支持 Get/Set 等基本操作

## 主要接口

```go
// 创建新的 LRU 缓存
func NewLruCache(capacity int) *LRUCache

// 获取缓存值
func (c *LRUCache) Get(key string) any error

// 存入缓存值
func (c *LRUCache) Set(key string, value any)
```

## 使用示例

```go
import "github.com/hacker4257/lrucache"

func main() {
    cache := lrucache.NewLruCache(2)
    cache.Put(1, "value1")
    value, found := cache.Get(1)
    if found {
        fmt.Println("Found:", value)
    }
}
```

## 实现原理

本项目使用双向链表 + 哈希表的经典实现方案:

1. 双向链表按照访问顺序存储节点,最近访问的节点位于表头
2. 哈希表存储 key 到链表节点的映射,实现 O(1) 查找
3. 当缓存满时,删除链表尾部节点(最久未使用)
4. 每次访问节点时,将其移动到链表头部

## 性能分析

- 时间复杂度: Get/Put 操作均为 O(1)
- 空间复杂度: O(capacity)

## 适用场景

- 需要缓存有限数量数据的场景
- 数据访问符合时间局部性原则的场景
- 内存敏感的应用

这个项目虽然实现简单,但体现了缓存系统的基本原理,可以作为学习缓存算法的良好示例
