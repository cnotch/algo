# algo
用 Go 实现基础数据结构和算法，尽量使之能用于生产环境。一些常用的算法和数据结构，Go的标准库是更好的选择；比如：排序，双向链表。

## 算法
### 排序
排序算法实现尽量保持和标准库排序接口一致。包括：
+ 冒泡排序([bubble](./sort/bubble/sort.go))
+ 选择排序([selectionn](./sort/selection/sort.go))
+ 插入排序([insertion](./sort/insertion/sort.go))
+ 希尔排序([shell](./sort/shell/sort.go))
+ 归并排序([merge](./sort/merge/sort.go))
+ 堆排序([heap](./sort/heap/sort.go))
+ 快速排序([quick](./sort/quick/sort.go))
+ 内省排序([intro](./sort/intro/sort.go))
+ 计数排序([counting](./sort/counting/sort.go))
+ 基数排序([radix](./sort/radix/sort.go))
+ 桶排序([bucket](./sort/bucket/slice.go))

### 顺序统计量
用于查找第 i 个顺序统计量（order statistic）是该集合中第 i 小的元素
+ 最大最小值([stat/minmax](./math/stat/minmax.go))
+ 选择第 i 小元素([stat/select](./math/stat/select.go))

### 数据结构
#### 基础
+ 栈([stack](./container/stack/stack.go))
+ 队列([queue](./container/queue/queue.go))
+ 单向链表([list/single](./container/list/single/list.go))
+ 双向链表([list/doubly](./container/list/doubly/list.go))
+ 跳跃链表([list/skip](./container/list/skip/list.go))
+ 无限制分支有根树([tree/unbounded](./container/tree/unbounded/tree.go))

#### 散列表

#### 二叉树

#### 前缀树
+ 朴素前缀树([trie/simple](./container/trie/simple/trie.go))
+ 压缩前缀树([trie/compress](./container/trie/compress/trie.go))
+ Critbit 前缀树([trie/critbit](./container/trie/critbit/trie.go))
+ 三分前缀树([trie/ternary](./container/trie/ternary/trie.go))
+ double-array 前缀树([trie/darts](./container/trie/darts/trie.go))