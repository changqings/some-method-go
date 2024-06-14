# some-method-go

一些与算法实现有关的方法:

### 栈与队列

var ss = []int{}

- 切片模拟栈(stack 先进后出)实现:
  push(),元素入栈，添加到栈顶; ss = append(ss,v)
  pop(),栈顶元素出栈; v = ss[len(ss)-1]; ss = ss[:len(ss)-1]
  peek(),访问栈顶元素, v = ss[len(ss)-1]
- 切片模拟队列(queue 先进先出)实现:
  push(),元素入栈，元素添加到队尾; ss = append(ss,v)
  pop(), 队首元素出列; v = ss[0], ss = ss[1:]
  peek(), 访问队首元素; v = ss[o]

### 排序

- sort.Ints() or slices.Sort()
- 使用插入排序
