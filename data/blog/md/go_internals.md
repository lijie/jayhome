## Go Internals

### Container

#### Slice

1. slice在内存中是连续的

2. append在slice的cap不够的情况下, 会分配新内存并copy老的slice到新内存
具体的实现在runtime/slice.go的growslice()
如果new_cap >= old_cap * 2 则之间增长到new_cap
else 如果old_cap<1024, 则反复double直到超过new_cap
else 反复增加old_cap/4直到超过new_cap

3. a := b[n:m] 实际上生成了一个新的slice 不过data是指向老的slice data
如果又执行a = append(a, ...) 则a可能指向老的, 也可能因为slice grow指向新的data

#### map

### Lock

#### Mutex

### goroutine

goroutine 相关的函数实现大部分在
runtime/proc1.go
另外, runtime/proc.go 应该是go编译后程序真正的入口

//生成新的goroutine并将其放到等待执行的队列中
func newproc1() {
  // 从local或者global gfreelist 里面分配一个g
  gfget()
  // new一个g
  malg()
  // 放到local或者global runqueue 等待运行
  runqput()
}

### reference

A Manual for the Plan 9 assembler
http://www.plan9.bell-labs.com/sys/doc/asm.html

