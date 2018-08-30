package main

import (
	"net"
	"net/rpc"
	"net/http"
	"fmt"
)

type Args struct {
	A, B int
}

//定义一个算术类型，其实就是int
type Arith int

//实现乘法的方法绑定到Arith类型，先不管为什么是这样的形式
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	//得到一个Arith类型的指针实例
	arith := new(Arith)
	//注册到rpc服务
	rpc.Register(arith)
	//挂到http服务上
	rpc.HandleHTTP()
	//开始监听
	fmt.Print("------------RPC服务端已经启动，现在可以开启客户端")
	l, _ := net.Listen("tcp", ":1234")
	http.Serve(l, nil)
}