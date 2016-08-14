//  求平方数
// 第一阶段是 gen 函数，它能够将一组整数转换为 channel ，
// channel 可以将数字发送出去。 gen 函数首先启动一个 goroutine ，
// 该 goroutine 发送数字到 channel ，当数字发送完时关闭 channel
func gen(nums ...int) <-chan int{
  out := make(chan int)
  go func ()  {
    for _, n := range nums {
      out <-n
    }
    close(out)  //当数字发送完时关闭 channel
  }()
  return out
}

 // sq 函数，它从 channel 接收一个整数，然后返回 一个 channel ，返回的 channel
 // 可以发送 接收到整数的平方。 当它的 inbound channel 关闭，并且把所有数字均发送到下游时，会关闭 outbound channel

 func sq(in <-chan int) <-chan int{
   out := make(chan int)
   go func() {
     for n:= range in{
       out <- n*n
     }
     close(out) // 当数字发送完时关闭 channel
   }()
   return out
 }

 //main 函数 用于设置流水线并运行最后一个阶段。最后一个阶段会从第二阶段接收数字，并逐个打印出来，
 // 直到来自于上游的 inbound channel 关闭

 func main() {
   // 设置流水线
   c := gen(2, 3)
   out := sq(c)

   // 消费输出结果
   fmt.Println(<-out) // 4
   fmt.Println(<-out) // 9
}

// 由于 sq 函数的 inbound channel 和 outbound channel 类型一样，
// 所以组合任意个 sq 函数。比如像下面这样使用：

func main() {
    // 设置流水线并消费输出结果
    for n := range sq(sq(gen(2, 3))) {
        fmt.Println(n) // 16 then 81
    }
}
