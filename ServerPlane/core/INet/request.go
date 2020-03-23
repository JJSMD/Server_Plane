package INet

import 	"../../core/IFace"

type Request struct {
	conn IFace.IConnection //已经和客户端建立好的 链接
	msg  IFace.IMessage    //客户端请求的数据
}

//获取请求连接信息
func(r *Request) GetConnection() IFace.IConnection {
	return r.conn
}
//获取请求消息的数据
func(r *Request) GetData() []byte {
	return r.msg.GetBody()
}

//获取请求的消息的ID
func (r *Request) GetMsgName() string {
	return r.msg.GetMsgName()
}