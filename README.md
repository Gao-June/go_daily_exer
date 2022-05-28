
## 2021.11.26  
>### 01 context 01  
    sync.WaitGroup 练习  

>### 02 context  
    sync.WaitGroup 练习  
>### 03 context  
    chan 练习

## 2021.11.27  
>### 01 kafka  
    基于sarama第三方库开发的kafka client（代码有问题）

>### 02 kafka consumer  
    kafka 练习（代码有问题）

>### 02 m  
    hello world  

## 2021.11.28
>### 01 load fun test  
>> Fun_sum  
    **Fun_sum.go**  
    自定义库  
**main.go**  
    主函数

>### 02 slice  
    slice 练习

## 2021.11.29
>### 01 slice  
    切片是 引用类型

>### 02 struct  
    结构体 声明
	推荐使用方法-3

>### 03 curl  
    书上的示例, http 简单使用

>### 04 interface  
    练习 接口

>### 05 goroutine  
    使用 goroutine 并发

>### 06 chan
    练习使用 管道 channel
	无缓冲 channel

>### 07 goroutine chan  
    并发在 无缓冲管道 读取值（任务）

>### 08 goroutine chans
    并发在 带缓冲管道 读取值（任务）


## 2021.11.30
>### no1_runner  
>>**main.go**  
    这个实例程序演示如何使用通道来监视
	程序运行的时间，以及程序运行时间过长时 如何终止程序
>>runner  
**runner.go**  
    runner 扩展包
	本项目用于展示 如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以用 runner 来终止

>### no2
>>**main.go**  
    函数调用练习
>>ppl  
**ppl.go**  
    自定义库

## 2021.12.11
>### 01 test  
    简单 test

## 2021.12.14  
>### 01 struct
    练习使用 struct 的几种常用方式
>### 02 error  
    练习使用 eerors.New 创造新的错误
>### 03 switch case  
    测试 switch case 中单条件为空的情况  
	单条件，内容为空、隐式 “ case x==10 : break; ”
>### 04 for range  
    无论是普通循环，还是 range迭代，其定义的局部变量都会重复使用

## 2021.12.15  
>### 01 customize library  
    写了一个调用 自定义库 的测试案例  
	能正常运行，不过不知道为啥 调用库包 Sum_func 会提示问题。  

## 2021.12.26  
>### 01 sync + go  
    写一个 go 协程， 用 sync 做同步

>### 02 channel with buffer  
    写一个有缓冲 channel

>### 03 channel without buffer
    写一个无缓冲的 channel

>### 04 channel without buffer  
    再试着写一次 goroutine + channel without buffer  
	上一个测试代码是用了 for 无限循环来时间 channel通信的，这次直接看看  
	可以，并不是要 for 无限循环！  

## 2022.05.08
>### 重新捡起 hello world
>### 使用 _ 可以表示匿名变量


## 2022.05.09
>### 练习 for 循环、数组、for-range、地址


## 2022.05.10
>### 练习使用多维数组
>### 练习切片、复制、遍历、添加元素
>### copy() 复制

## 2022.05.11
>### golang 中也能实现  C++里的函数指针
>### 练习结构体

## 2022.05.12
>### 方法和接收者
>### scan 读取输入
>### 接口
>### 包的导入
>### 包的导入
>### 读取文件

## 2022.05.13
>### 练习调用 自定义库
>### 练习使用 goroutine, 及 sync.WaitGroup 
>### 练习管道 channel
>### 用 for range 循环读取 channel里的数

## 2022.05.23
>### 在抖音看到一个有意思的线程时间问题（很有意思）
    我在这里分别用 c++、golang、python试了下
>### go语言还支持中文作为 变量名！
>### iota 默认从0开始计数，可以从自定义的位置开始计数吗？
>### 探究下 golang 对类型转换的支持（隐式转换、显示转换）

## 2020.05.24
>### for-range
>### 探究 defer 和 return 的顺序

## 2022.05.25
>### 练习 单元测试
>### 《项目》01 - 01 - 网络编程 TCP

## 2022.05.26
>### 练习 func() strings.Trim
>### 《项目》01-02 - 网络编程 UDP
>### 《项目》01-03 - http编程
>### 《项目》01-04 - socket编程
>### 练习 goroutine
>### 如果主协程退出了，其他任务还执行吗（代码测试）
>### runtime.Gosched()  允许先运行其它 协程
>### runtime.Goexit()   退出当前协程
>### runtime.Goexit()-2 再写一个退出当前协程
>### 练习使用 tuntime.GOMAXPROCS
>### 无缓冲 chan 
>### 有缓冲 chan
>### 这里写一个 goroutine 实例， 同时练习一下 chan
		本质上是生产者消费者模型
## 2022.05.27
>###  练习使用 time 时间类型
>###  爬虫测试，有点意思
>### 《Gin框架》 01 - 01- gin路由 初试
>### 《Gin框架》 01 - 03 - gin路由 通过Context的Param方法来获取API参数
>### 《Gin框架》 01 - 04 - gin路由 URL 参数可以通过 DefaultQuery() 或 Query() 方法获取
>### 《Gin框架》 01 - 05 - gin路由 表单传输为 post 请求
>### 《Gin框架》 01 - 06 - gin路由 上传单个文件