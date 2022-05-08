
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
    