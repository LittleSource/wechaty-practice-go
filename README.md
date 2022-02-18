# wecahty golang版本SDK实践
简单记录一下过程

1. 创建项目
2. 参考<https://github.com/wechaty/go-wechaty>的README
3. 启动项目
   
   很好，启动报错了，没有token
4. 获取token<http://120.55.60.194/#/login>
5. 将token设置到环境变量后启动成功
6. 在接收消息事件中去回复消息
   ```go
   wechatyServer.OnMessage(func(context *wechaty.Context, message *user.Message) {
    _,err := message.Say("ok")
    if err != nil{
        fmt.Printf("error: %s\n", err.Error())
    }
    fmt.Printf("Message: %s\n", message)
   })
   ```
> 用自己的账号测试的，结果某个群正在聊天收到了我发的十几个“ok”。。。