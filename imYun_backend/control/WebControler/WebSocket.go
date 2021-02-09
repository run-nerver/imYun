package WebControler

import (
	"PrintYun/libs"
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"strconv"
	"time"
)

const namespace = "default"

var Ns *neffos.Server
//var ToSocket chan interface{}
var NsMessage = websocket.Message{
	Body: []byte("Hello world"),
	IsNative: true,
}

func SocketServer() *neffos.Server {
	events := make(neffos.Namespaces)

	events.On(namespace,"Authorization", func(ns *neffos.NSConn, msg neffos.Message) error {
		data := string(msg.Body)
		fmt.Println(data)
		CLaim, err := libs.ParseHStoken(data)
		if err != nil {
			ns.Emit("notify",[]byte("1002"))
			ns.Conn.Close()
			return err
		}
		RDB := libs.GetRedisDB2()
		_, err = RDB.HSet(libs.Ctx, strconv.Itoa(int(CLaim["id"].(float64))), map[string]interface{}{"JWT":data, "WSID":ns.Conn.ID()}).Result()
		if err != nil{
			ns.Emit("notify",[]byte("1002"))
			ns.Conn.Close()
			return err
		}
		return neffos.Reply([]byte("Hello : " + CLaim["username"].(string)))
	})
	Ns = websocket.New(websocket.DefaultGobwasUpgrader, events)
	//Ns.OnConnect = func(ns *neffos.Conn) error {
	//	RDB := libs.GetRedisDB2()
	//	_, err := RDB.HSet(libs.Ctx, strconv.Itoa(int(CLaim["id"].(float64))), map[string]interface{}{"WSID":ns.Conn.ID()}).Result()
	//	if err != nil{
	//		ns.Emit("notify",[]byte("1002"))
	//		ns.Conn.Close()
	//		return err
	//	}
	//	return nil
	//}
	fmt.Println(Ns.GetConnections())

	time.AfterFunc(time.Second*20, func() {
		Ns.Broadcast(nil, neffos.Message{
			Namespace: namespace,
			Event:     "Notify",
			Body:      []byte("Ping!"),
		})
	})

	return Ns
}

func IdGen(ctx *context.Context) string {
	var (
		UserID string
	)

	CLaim, _ := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))

	UserID = strconv.Itoa(int(CLaim["id"].(float64)))
	if UserID != "" {
		return UserID
	}
	return websocket.DefaultIDGenerator(ctx)
}

func GetNs() *neffos.Server {
	return Ns
}



















//var Events = websocket.Namespaces{
//	namespace : websocket.Events{
//		websocket.OnNamespaceConnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
//			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
//			ctx := websocket.GetContext(nsConn.Conn)
//
//			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
//				nsConn, msg.Namespace,
//				ctx.RemoteAddr())
//			return nil
//		},
//
//		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
//			log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
//			return nil
//		},
//
//		"chat" : func(nsConn *neffos.NSConn, msg neffos.Message) error {
//			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
//			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))
//
//			// Write message back to the client message owner with:
//			// nsConn.Emit("chat", msg)
//			// Write message to all except this client with:
//			nsConn.Conn.Server().Broadcast(nsConn, NsMessage)
//			nsConn.Conn.Write(NsMessage)
//			return nil
//		},
//	},
//}
//var Events = websocket.Events{
//	websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
//		log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
//		nsConn.Conn.Server().Broadcast(nsConn, msg)
//		return nil
//	},
//	"OnConnect" : func(conn *neffos.NSConn, message neffos.Message) error {
//		log.Printf("[%s] Connected to server!", c.ID())
//		return nil
//	},
//	"OnDisconnect" : func(conn *neffos.NSConn, message neffos.Message) error {
//		log.Printf("[%s] Disconnected from server", c.ID())
//		return nil
//	},
//}

// 如果启动 yaag 需要关闭 否则在 yaag 哪里会进行报错
//func WebSocketServer() *neffos.Server {
	//var Events = websocket.Events{
	//	websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
	//		log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
	//		nsConn.Conn.Server().Broadcast(nsConn, msg)
	//		return nil
	//	},
	//}
	//Ns = websocket.New(websocket.DefaultGobwasUpgrader, Events)
	//Ns.OnConnect = func(c *neffos.Conn) error {
	//	log.Printf("[%s] Connected to server!", c.ID())
	//	return nil
	//}
	//

	//Ns.OnDisconnect = func(c *neffos.Conn) {
	//	log.Printf("[%s] Disconnected from server", c.ID())
	//}
	//return Ns
//}

//func GetNs() *neffos.Server {
//	return Ns
//}
//
//
//func OnNamespaceMessage(nsConn *websocket.NSConn, msg websocket.Message) error {
//	log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
//	nsConn.Conn.Write(msg)
//	return nil
//}
//
//func OnNamespaceConnected(nsConn *websocket.NSConn, msg websocket.Message) error {
//	ctx := websocket.GetContext(nsConn.Conn)
//
//	nsConn.Conn.Write(NsMessage)
//	log.Printf("[%s] connected to namespace [%s] with IP [%s]",
//		nsConn, msg.Namespace,
//		ctx.GetID())
//	return nil
//}
//
//
//func OnNamespaceDisconnect(nsConn *websocket.NSConn, msg websocket.Message) error {
//
//	nsConn.Conn.Write(NsMessage)
//	log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
//	return nil
//}

//type IdGen func(*context.Context) string

//func IdGen(ctx *context.Context) string {
//	var (
//		UserID string
//	)
//	UserID = "1"
//	if UserID == "1" {
//		return UserID
//	}
//	return websocket.DefaultIDGenerator(ctx)
//}



//
//func OnNamespaceConnected(nsConn *websocket.NSConn, msg websocket.Message) error {
//	// with `websocket.GetContext` you can retrieve the Iris' `Context`.
//	ctx := websocket.GetContext(nsConn.Conn)
//
//	log.Printf("[%s] connected to namespace [%s] with IP [%s]",
//		nsConn,
//		msg.Namespace,
//		ctx.RemoteAddr())
//	return nil
//}
//
//func OnNamespaceDisconnect(nsConn *websocket.NSConn, msg websocket.Message) error {
//	log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
//	return nil
//}
//
//func Chat(nsConn *websocket.NSConn, msg websocket.Message) error {
//	// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
//	log.Printf("[%s] sent: %s", nsConn, string(msg.Body))
//
//	// Write message back to the client message owner with:
//	// nsConn.Emit("chat", msg)
//	// Write message to all except this client with:
//	nsConn.Conn.Server().Broadcast(nsConn, msg)
//	return nil
//}