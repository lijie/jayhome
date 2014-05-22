## Redis学习笔记

### 存储结构

#### String

简单的kv存储，除了set/get之外还支持append，setbit，inc等操作

#### List

kv的扩展，v是一个链表，链表的常见操作都可以使用

#### Hash

kv的扩展，v是一个字典，类似mongodb的文档

#### Set

kv的扩展，v是一个类似C++中的set，唯一且可快速查找

#### Sorted Set

kv的扩展，v是一个排序的set，排序的key是额外指定的一个score。
内部实现上，set还是使用的dict，并额外将数据插入一份到一个skiplist实现排序

### 内部存储结构

#### Dict

redis最重要的数据结构，是一个可自动扩展的hash-table
每一个dict实际上最多持有2个hashtable，大部分情况下使用的是0
发生hash扩展时，会new一个1，然后逐步将0中的entry迁移到1中，
完成迁移后，释放掉旧的，把新的赋值给0

dict的迁移不是异步的，而是一旦需要迁移时，每次访问该dict，
就完成一个slot的迁移。

redis中，几乎所有支持快速查询的数据都是用dict来存储的。

#### ziplist

一个经过压缩的双向链表

#### zsklist

跳跃表的一种实现，查询修改的效率接近红黑树。
redis中的实现有些特别，比如允许key重复。

zsklist在redis中用来给sorted set排序。

### 客户端通信

#### 协议

redis的协议有一个设计要点是可读，所以它可以说是一种文本协议。
不过可以传输二进制数据.

* 如何传递二进制数据?

#### pub/sub模型

### 部署和容灾

#### 数据冗余

redis支持Master-Slave模型

* 最多多少个Slave?
* 同步延时?
* Master写成功就算成功?还是需要所有或者部分Slave同步成功才算成功?

#### 数据持久化

redis内置2种数据持久化方案, RDB和AOF

RDB: redis会周期性的将整个db数据dump到本地的一个文件.
RDB的完成过程是由redis fork出来的一个子进程来做, 并不会影响redis主进程本身.

RDB因为是fork子进程来处理所有数据, 所以它存在一个可能性, 就是在dump数据的过程中, 内存使用量可能会double.

AOF: redis将所有的写操作的请求协议append一个log文件中, redis重启时会replay该文件的数据.

RDB与AOF可以结合使用.

#### 数据分割

在单个实例无法保存所有数据,或者为了提高CPU利用率等情况下, 需要数据分割到多个实例.

分割的方案一般是range或者hash, hash又有简单的取模或者一致性hash等方案.
当redis仅仅作为cache使用时,一致性hash就已经足够,如果redis作为最终数据存储,那就需要完善的cluster方案.
主要是需要在多个node之间做数据的rebalance.

redis目前内置的partioning方案redis cluster还处于实验性质.

redis目前推荐的方案是使用外部组件Twemproxy.

redis的partitioning是基于key的, 如果某个key特别大,
大到一台机器无法存储, 这对于现有的redis partitiong方案来说都无法解决.

在cluster情况下,某些redis功能会失效, 主要是事务或者pipelining功能, 在处理多个key时,
如果这些key分布在不同机器上, 就会失败.

redis cluster跟Twemproxy不一样的是, 它没有引入一个proxy,
而是允许客户连接任意一个cluster中的redis实例, 如果某个client请求的数据不在当前实例,
cluster会返回一个redirect, 告诉客户端正确的地址.


### 高级应用

#### LRU Cache

#### 分布式锁

### Other Topics

#### 内存优化
