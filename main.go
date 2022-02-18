package main


import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

func main() {
	wechatyServer := wechaty.NewWechaty()
	wechatyServer.OnScan(func(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		fmt.Printf("Scan QR Code to login: %s\nhttps://wechaty.github.io/qrcode/%s\n", status, qrCode)
	})
	wechatyServer.OnLogin(func(context *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s 已登录\n", user)
	})
	wechatyServer.OnMessage(func(context *wechaty.Context, message *user.Message) {
		_,err := message.Say("ok")
		if err != nil{
			fmt.Printf("error: %s\n", err.Error())
		}
		fmt.Printf("Message: %s\n", message)
	})
	wechatyServer.DaemonStart()
}
