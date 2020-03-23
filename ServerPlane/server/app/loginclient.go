package app

import (
	"../../core/IFace"
	"../../core/INet"
	"../../proto"
	"../../utils"
	"encoding/json"
	"errors"
	"time"
)

var (
	lPingTimerId uint32
)

func lClientConnStart(conn IFace.IConnection){
	utils.Log.Info("lClientConnStart:%s", conn.RemoteAddr().String())
	lPingTimerId, _ = utils.Scheduler.NewTimerInterval(15*time.Second,utils.IntervalForever, loginPingTimer, []interface{}{conn})
}

func lClientConnStop(conn IFace.IConnection){

	client := conn.GetTcpNetWork().(*INet.Client)
	Name := client.GetName()
	Id := client.GetId()
	IP := client.GetHost()
	Port := client.GetPort()
	cType := client.GetClientType()

	utils.Log.Info("lClientConnStop:%s,%s,%s:%d,%d", Name, Id, IP, Port, cType)
	utils.Scheduler.NewTimerAfter(5*time.Second, restartLoginClient, []interface{}{Name, Id, IP, Port, cType})
	utils.Log.Info("lClientConnStop end")

	if lPingTimerId > 0{
		utils.Scheduler.CancelTimer(lPingTimerId)
		lPingTimerId = 0
	}
}

func LoginClient(clientName string, clientId string,
	remoteHost string, remotePort int, clientType proto.ServerType) (*INet.Client, error){

	if remotePort > 0 && remoteHost != ""{
		var c *INet.Client
		c = INet.NewClient(clientName, clientId, remoteHost, remotePort)
		c.SetClientType(clientType)
		c.SetOnConnStart(lClientConnStart)
		c.SetOnConnStop(lClientConnStop)
		c.Running()
		return c, nil
	}
	return nil, errors.New("new LoginClient Error")
}


func restartLoginClient(v ...interface{}) {
	Name := v[0].(string)
	Id := v[1].(string)
	Ip := v[2].(string)
	Port := v[3].(int)
	cType := v[4].(proto.ServerType)

	LoginClient(Name, Id, Ip, Port, cType)
}

func loginPingTimer(v ...interface{})  {

	conn := v[0].(IFace.IConnection)
	info := proto.PingPong{}
	info.CurTime = time.Now().Unix()

	data ,err := json.Marshal(info)
	if err == nil{
		conn.SendMsg(proto.SystemPing, data)
	}else{
		utils.Log.Info("loginPingTimer error:%s", err.Error())
	}
}

