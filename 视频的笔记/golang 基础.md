# golang 基础

#### 4.1 map的基本使用

##### 4.1.1.1 声明 & 默认值

```go
// 声明
var m map[string]string
```

map的声明的时候默认值是**nil** ，此时进行取值，返回的是**对应类型的零值**（不存在也是返回零值）。

例子：

```
// bool 的零值是false
var m map[int]bool 
a, ok := m[1]
fmt.Println(a, ok) // false  false
1234
```

// int 的零值是0
var m map[int]int
a, ok := m[1]
fmt.Println(a, ok) // 0 false



##### 4.1.1.2 初始化

```
// 声明之后必须初始化，才能使用
m = make(map[string]int)
m = map[string]int{}
```

// 声明并初始化 <= 注意这里是 := 不是 =
m := make(map[string]int)
m := map[string]int{1:1}

**向未初始化的map赋值引起 panic:** assign to entry in nil map.

##### 4.1.1.3 key与value的限制

key一定要是**可比较**的类型（可以理解为支持==的操作）：

| 可比较类型                  | 不可比较类型 |
| --------------------------- | ------------ |
| boolean                     | slice        |
| numeric                     | map          |
| string                      | func         |
| pointer                     |              |
| channel                     |              |
| interface                   |              |
| 包含前文类型的array和struct |              |

如果是非法的key类型，会报错：invalid map key type xxx。

> golang为uint32、uint64、string提供了fast access，使用这些类型作为key可以提高map访问速度。[runtime/hashmap_fast.go]

value可以是**任意类型**。

##### 4.1.1.4 新增 & 删除 & 更新 & 查询

```go
// 新增
m["name"] = "咖啡色的羊驼"
```

// 删除，key不存在则啥也不干
delete(m, “name”)

// 更新
m[“name”] = “咖啡色的羊驼2”

// 查询，key不存在返回value类型的零值
i := m[“name”] // 三种查询方式，
i, ok := m[“name”]
_, ok := m[“name”]



##### 4.1.1.5 遍历

需要强调的是map本身是**无序的**，在遍历的时候并不会按照你传入的顺序，进行传出。

正常遍历：

```go
for k, v := range m { 
    fmt.Println(k, v)
}
```

有序遍历：

```go
import "sort"
var keys []string
// 把key单独抽取出来，放在数组中
for k, _ := range m {
    keys = append(keys, k)
}
// 进行数组的排序
sort.Strings(keys)
// 遍历数组就是有序的了
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

##### 4.1.1.6 函数传参

Golang中是没有引用传递的，均为值传递。这意味着传递的是数据的拷贝。
那么map本身是**引用类型**，作为形参或返回参数的时候，传递的是**地址的拷贝**，**扩容**时也**不会改变**这个地址。

```go
var m map[int64]int64
m = make(map[int64]int64, 1)
fmt.Printf("m 原始的地址是：%p\n", m)
changeM(m)
fmt.Printf("m 改变后地址是：%p\n", m)
fmt.Println("m 长度是", len(m))
fmt.Println("m 参数是", m)
```

// 改变map的函数
func changeM(m map[int64]int64) {
fmt.Printf(“m 函数开始时地址是：%p\n”, m)
var max = 5
for i := 0; i < max; i++ {
m[int64(i)] = 2
}
fmt.Printf(“m 在函数返回前地址是：%p\n”, m)
}



输出：

```go
m 原始地址是：0xc42007a180
m 函数开始时地址是：0xc42007a180
m 在函数返回前地址是：0xc42007a180
m 改变后地址是：0xc42007a180
m 长度是 5
m 参数是  map[3:2 4:2 0:2 1:2 2:2]
```

#### 4.1.2 map的深入了解

##### 4.1.2.1 map的基础数据结构 & 图

```go
type hmap struct {
    count        int  //元素个数
    flags        uint8   
    B            uint8 //扩容常量
    noverflow    uint16 //溢出 bucket 个数
    hash0        uint32 //hash 种子
    buckets      unsafe.Pointer //bucket 数组指针
    oldbuckets   unsafe.Pointer //扩容时旧的buckets 数组指针
    nevacuate    uintptr  //扩容搬迁进度
    extra        *mapextra //记录溢出相关
}
```

type bmap struct {
tophash [bucketCnt]uint8
// Followed by bucketCnt keys
//and then bucketan Cnt values
// Followed by overflow pointer.
}

说明：**每个map的地层结构是hmap，是有若干个机构为bmap的bucket组成的数组，每个bucket可以存放若干个元素(通常是8个)，那么每个key会根据hash算法归到同一个bucket中，当一个bucket中的元素超过8个的时候，hmap会使用extra中的overflow来扩展存储key**。

来一个图，方便记忆：

![map](D:/go/goProject/src/go_code/视频的笔记/img/map.png)

##### 4.1.2.2 map的hash值计算

那么具体key是分配到哪个bucket呢？也就是bmap中的tophash是如何计算？

**golang为每个类型定义了类型描述器_type，并实现了hashable类型的_type.alg.hash和_type.alg.equal**

```go
type typeAlg struct {
    // function for hashing objects of this type
    // (ptr to object, seed) -> hash
    hash func(unsafe.Pointer, uintptr) uintptr
    // function for comparing objects of this type
    // (ptr to object A, ptr to object B) -> ==?
    equal func(unsafe.Pointer, unsafe.Pointer) bool
```

具体实现文件：go/1.10.3/libexec/src/runtime/hashmap.go:

```go
// tophash calculates the tophash value for hash.
func tophash(hash uintptr) uint8 {
	top := uint8(hash >> (sys.PtrSize*8 - 8))
	if top < minTopHash {
		top += minTopHash
	}
	return top
}
```

##### 4.1.2.3 map的查找

具体实现文件：go/1.10.3/libexec/src/runtime/hashmap.go:

```go
// mapaccess1 returns a pointer to h[key].  Never returns nil, instead // it will return a reference to the zero object for the value type if // the key is not in the map. // NOTE: The returned pointer may keep the whole map live, so don't // hold onto it for very long. func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {    ...    // 并发访问检查    if h.flags&hashWriting != 0 {        throw("concurrent map read and map write")    } 
// 计算key的hash值
alg := t.key.alg
hash := alg.hash(key, uintptr(h.hash0)) // alg.hash

// hash值对m取余数得到对应的bucket
m := uintptr(1)&lt;&lt;h.B - 1
b := (*bmap)(add(h.buckets, (hash&amp;m)*uintptr(t.bucketsize)))

// 如果老的bucket还没有迁移，则在老的bucket里面找
if c := h.oldbuckets; c != nil {
    if !h.sameSizeGrow() {
        m &gt;&gt;= 1
    }
    oldb := (*bmap)(add(c, (hash&amp;m)*uintptr(t.bucketsize)))
    if !evacuated(oldb) {
        b = oldb
    }
}

// 计算tophash，取高8位
top := uint8(hash &gt;&gt; (sys.PtrSize*8 - 8))

for {
    for i := uintptr(0); i &lt; bucketCnt; i++ {
        // 检查top值，如高8位不一样就找下一个
        if b.tophash[i] != top {
            continue
        }
        
        // 取key的地址
        k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
        
        if alg.equal(key, k) { // alg.equal
            // 取value得地址
            v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
        }
    }
   
    // 如果当前bucket没有找到，则找bucket链的下一个bucket
    b = b.overflow(t)
    if b == nil {
        // 返回零值
        return unsafe.Pointer(&amp;zeroVal[0])
    }
}
```



##### 4.1.2.4 map的更新/插入过程

```go
// Like mapaccess, but allocates a slot for the key if it is not present in the map. func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer { 
// 如果已经达到了load factor的最大值，那我们就继续开始扩容
if !h.growing() &amp;&amp; (overLoadFactor(int64(h.count), h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
    hashGrow(t, h)
    goto again 
}

if inserti == nil {
    // burrent满了的话，需要申请一个新的
    newb := h.newoverflow(t, b)
    inserti = &amp;newb.tophash[0]
    insertk = add(unsafe.Pointer(newb), dataOffset)
    val = add(insertk, bucketCnt*uintptr(t.keysize))
}

// 在插入的位置，存储键值
if t.indirectkey {
    kmem := newobject(t.key)
    *(*unsafe.Pointer)(insertk) = kmem
    insertk = kmem
}
if t.indirectvalue {
    vmem := newobject(t.elem)
    *(*unsafe.Pointer)(val) = vmem
}
typedmemmove(t.key, insertk, key)
*inserti = top
h.count++
}
```



##### 4.1.2.5 map的删除过程

```go
func mapdelete(t *maptype, h *hmap, key unsafe.Pointer) {
    // 找key
   。。。
    // 若找到把对应的tophash里面的打上空的标记
    b.tophash[i] = empty
    h.count--
}
```

##### 4.1.2.6 map的扩容机制

map判断扩容的函数：

```go
// overLoadFactor reports whether count items placed in 1<<B buckets is over loadFactor.
func overLoadFactor(count int, B uint8) bool {
    // 注意这里有一个loadFactorNum/loadFactorDen
	return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
}
```

func bucketShift(b uint8) uintptr {
if sys.GoarchAmd64|sys.GoarchAmd64p32|sys.Goarch386 != 0 {
b &= sys.PtrSize*8 - 1 // help x86 archs remove shift overflow checks
}
return uintptr(1) << b
}



每次map进行更新或者新增的时候，会先通过以上函数判断一下load factor。来决定是否扩容。

> 扩容白话文：如果之前为2^n ，那么下一次扩容是2^(n+1),每次扩容都是之前的两倍。扩容后需要重新计算每一项在hash中的位置，新表为老的两倍，此时前文的oldbacket用上了，用来存同时存在的两个心就map，等数据迁移完毕就可以释放oldbacket了

好处：**均摊扩容时间，一定程度上缩短了扩容时间**（是不是和gc的引用计数法有点像，都是均摊～）

那么overLoadFactor函数中有一个常量6.5（loadFactorNum/loadFactorDen）来进行影响扩容时机。这个值的来源是测试取中的结果。

```
  loadFactor    %overflow  bytes/entry     hitprobe    missprobe
        4.00         2.13        20.77         3.00         4.00
        4.50         4.05        17.30         3.25         4.50
        5.00         6.85        14.77         3.50         5.00
        5.50        10.55        12.94         3.75         5.50
        6.00        15.27        11.67         4.00         6.00
        6.50        20.90        10.79         4.25         6.50
        7.00        27.14        10.15         4.50         7.00
        7.50        34.03         9.73         4.75         7.50
        8.00        41.10         9.40         5.00         8.00
```

| 字段        | 说明                                         |
| ----------- | -------------------------------------------- |
| %overflow   | 溢出率，平均一个bucket有多少个kv的时候会溢出 |
| bytes/entry | 平均存一个kv需要额外存储多少字节的数据       |
| hitprobe    | 找到一个存在的key平均需要找几下              |
| missprobe   | 找到一个不存在的key平均需要找几下            |

#### 4.1.3 并发中的map

##### 4.1.3.1 安全性

这里呢，实现一个小功能来证明下并发**不是安全**的。
并发起两个goroutine，分别对map进行数据的增加

```go
func main() { test := map[int]int {1:1} go func() { 	i := 0 	for i < 10000 { 		test[1]=1 		i++ 	} }() 
go func() {
	i := 0
	for i &lt; 10000 {
		test[1]=1
		i++
	}
}()

time.Sleep(2*time.Second)
fmt.Println(test)
}
```

会发现有这样的报错：

```
fatal error: concurrent map read and map write
```

根本原因就是：并发的去读写map结构的数据了。

##### 4.1.3.2 处理方案 & 优缺点

那解决方案就是加锁。上代码

```go
func main() { test := map[int]int {1:1} var s sync.RWMutex go func() { 	i := 0 	for i < 10000 { 		s.Lock() 		test[1]=1 		s.Unlock() 		i++ 	} }() 
go func() {
	i := 0
	for i &lt; 10000 {
		s.Lock()
		test[1]=1
		s.Unlock()
		i++
	}
}()

time.Sleep(2*time.Second)
fmt.Println(test)
}
```

> 优点：实现简单粗暴，好理解
> 缺点：锁的粒度为整个map，存在优化空间
> 适用场景：all

##### 4.1.3.3 官方处理方案 & 优缺点

想一想，在程序设计中，想增加运行的速度，那么必然要有另外的牺牲，很容易想到“空间换时间”的方案，现在来实战体验一把。

```go
func main() { test := sync.Map{} test.Store(1, 1) go func() { 	i := 0 	for i < 10000 { 		test.Store(1, 1) 		i++ 	} }() 
go func() {
	i := 0
	for i &lt; 10000 {
		test.Store(1, 1)
		i++
	}
}()

time.Sleep(time.Second)
fmt.Println(test.Load(1))
}
```

运行完呢，会发现，其实是不会报错的。因为sync.Map里头已经实现了一套加锁的机制，让你更方便地使用map。

sync.Map的原理介绍：***sync.Map里头有两个map一个是专门用于读的read map，另一个是才是提供读写的dirty map；优先读read map，若不存在则加锁穿透读dirty map，同时记录一个未从read map读到的计数，当计数到达一定值，就将read map用dirty map进行覆盖。***

> 优点：是官方出的，是亲儿子；通过空间换时间的方式；读写分离；
> 缺点：不适用于大量写的场景，这样会导致read map读不到数据而进一步加锁读取，同时dirty map也会一直晋升为read map，整体性能较差。
> 适用场景：大量读，少量写

##### 4.1.3.4 额外的处理机制

想一想，mysql加锁，是不是有表级锁、行级锁，前文的sync.RWMutex加锁方式相当于表级锁。

而sync.Map其实也是相当于表级锁，只不过多读写分了两个map，本质还是一样的。

既然这样，那就自然知道优化方向了：**就是把锁的粒度尽可能降低来提高运行速度**。

思路：对一个大map进行hash，其内部是n个小map，根据key来来hash确定在具体的那个小map中，这样加锁的粒度就变成1/n了。
网上找了下，真有大佬实现了：[点这里](https://github.com/orcaman/concurrent-map)

#### 4.1.4 map的gc回收机制

##### 4.1.4.1 实战代码 && 处理机制

我们知道呢，map在golang里头是只增不减的一种数组结构，他只会在删除的时候进行打标记说明该内存空间已经empty了，不会回收的。

show my code，no bb:

```go
var intMap map[int]int
```

func main() {
printMemStats(“初始化”)

```go
// 添加1w个map值
intMap = make(map[int]int, 10000)
for i := 0; i &lt; 10000; i++ {
	intMap[i] = i
}

// 手动进行gc操作
runtime.GC()
// 再次查看数据
printMemStats("增加map数据后")

log.Println("删除前数组长度：", len(intMap))
for i := 0; i &lt; 10000; i++ {
	delete(intMap, i)
}
log.Println("删除后数组长度：", len(intMap))

// 再次进行手动GC回收
runtime.GC()
printMemStats("删除map数据后")

// 设置为nil进行回收
intMap = nil
runtime.GC()
printMemStats("设置为nil后")
}
```



func printMemStats(mag string) {
var m runtime.MemStats
runtime.ReadMemStats(&m)
log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}

会输出：

```
初始化：分配的内存 = 65KB, GC的次数 = 0
增加map数据后：分配的内存 = 381KB, GC的次数 = 1
删除前数组长度： 10000
删除后数组长度： 0
删除map数据后：分配的内存 = 381KB, GC的次数 = 2
设置为nil后：分配的内存 = 68KB, GC的次数 = 3
```

很明显可以看到delete是不会真正的把map释放的，所以要**回收map还是需要设为nil**