
## [list](https://github.com/zhaojiaming110/go_data_structures/tree/master/list)

列表作为一种抽象数据类型，列表对象应支持以下操作接口。

| 操作接口                        | 功能                                                         | 适用对象 |
| :------------------------------ | :----------------------------------------------------------- | :------- |
| Size()                          | 报告列表当前的规模（节点总数）                               | 列表     |
| First()、Last()                 | 返回首、末节点的位置                                         | 列表     |
| InsertAsFirst()、InsertAsLast() | 将元素当作首、末节点插入                                     | 列表     |
| InsertBefore() 、InsertAfter()  | 将元素e当作节点p的直接前驱、后继插入                         | 列表     |
| Remove(p *node)                 | 删除节点p                                                    | 列表     |
| Disordered()                    | 判断所有节点是否已按非降序排列                               | 列表     |
| Find(e)                         | 查找目标元素e，失败时返回nil                                 | 列表     |
| Deduplicate()                   | 剔除重复节点                                                 | 列表     |
| Uniquify()                      | 剔除重复节点                                                 | 有序列表 |
| Sort()                          | 调整各节点的位置，使之按非降序排列<br/>支持插入排序、归并排序、选择排序 | 列表     |
| Traverse()                      | 遍历并统一处理所有节点<br/>处理方法由函数对象指定            | 列表     |
| Clear()                         | 删除列表                                                     | 列表     |
| Get()                           | 根据列表位置，返回目标节点元素                               | 列表     |
| Values()                        | 以切片的形式返回列表元素                                     | 列表     |

### Interface

```go
// Interface of List that all lists implements.
type List interface {
	Size() int
	Values() []interface{}
	First() *node
	Last() *node
	Get(index int) (interface{}, bool)
	InsertAsFirst(interface{})
	InsertAsLast(values ...interface{})
	InsertBefore(p *node, e interface{}) *node
	InsertAfter(p *node, e interface{})	*node
	Remove(p *node) interface{}
	Clear()
	Disordered() bool
	Sort(m int, p *node, n int)
	Find(e interface{}) *node
	Deduplicate() int
	Uniquify() int
	Traverse(visit func(interface{}))
}
```

 具体实现请看[Github](https://github.com/zhaojiaming110/go_data_structures/tree/master/list).

## [Stack](https://github.com/zhaojiaming110/go_data_structures/tree/master/Stacks)

栈（Stack）是存放数据对象的一种特殊容器，其中的数据元素按线性的逻辑次序排列。尽管栈结构也支持对象的插入和删除操作，但其操作的范围仅限于栈的某一定端。也就是说，若约定新元素只能从某一端插入其中，则反过来也只能从这一端删除已有的元素。禁止操作的另一端则称为盲端。

作为抽象数据类型，栈所支持的操作接口可归纳为下表：

| 操作接口 |        功    能        |
| :------: | :--------------------: |
|  Size()  |      报告栈的规模      |
| Empty()  |      判断栈否为空      |
| Push(e)  |      将e插入栈顶       |
|  top()   |      删除栈顶元素      |
|  Peek()  | 返回栈顶元素，但不删除 |

### Interface

```go
// Interface Stack that all stacks implement
type Stack interface {
	Size() int
	Empty() bool
	Push(value interface{}) bool
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	Clear()
	Values() (values []interface{})
}
```

实现栈可通过数组和单链表来实现。具体代码可见[Github](https://github.com/zhaojiaming110/go_data_structures/tree/master/Stacks).

### 栈的典型应用场合

- 逆序输出

    Conversion

    输出次序与处理过程颠倒；递归深度和输出长度不易预知

- 递归嵌套

    Stack permutation + parenthesis

    具有自相似性的问题可递归描述，但分支位置和嵌套深度不固定

- 延迟缓冲

    evaluation

    线性扫描算法模式中，在预读足够长之后，方能确定可处理的前缀

- 栈式计算

    RPN

    基于栈结构的特定计算模式

## [Queue](https://github.com/zhaojiaming110/go_data_structures/tree/master/queue)

与栈一样，队列也是存放数据对象的一种容器，其中的数据对象也按线性的逻辑次序排列。队列结构同样支持对象的插入和删除，但两种操作的范围分别被限制于队列的两端。 -----若约定新对象只能从某一端插入其中，则只能从另一端删除已有的元素。允许取出元素的一端陈作队头（front），而允许插入元素的另一端称为队尾。

作为一种抽象数据类型，队列结构必须支持以下操作接口。

|    操作    |            功能            |
| :--------: | :------------------------: |
|   Size()   | 报告队列的规模（元素总数） |
|  Empty()   |      判断队列是否为空      |
| Enqueue(e) |        将e插入队尾         |
| Dequeue()  |        删除队首对象        |
|  Front()   |        返回队首对象        |

### Interface

```go
type Queue interface {
	Size() int
	Empty() bool
	Enqueue(value interface{}) bool
	Dequeue() (value interface{}, ok bool)
	Front()	(value interface{}, ok bool)
	Values() (values []interface{})
	Clear()
}
```

队列实现也可以使用数组和链表的方式实现。值得注意的是用数组实现限制了列表的长度，且插入是在固定存储单位上的顺序循环插入，同样的删除操作也是在固定存储单位上的顺序循环删除。

详细实现，请参见[Github](https://github.com/zhaojiaming110/go_data_structures/tree/master/queue).

## [Tree](https://github.com/zhaojiaming110/go_data_structures/tree/master/trees/binaryTree)

### Interface

```go
type BinaryTree interface {
	InsertAsRoot(data interface{})
	InsertAsLC(x *binNode, data interface{})
	InsertAsRC(x *binNode, data interface{})
	Traverse(model int, visit func( interface{}))
}
```

