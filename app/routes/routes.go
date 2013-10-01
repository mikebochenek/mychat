// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tApplication struct {}
var Application tApplication


func (p tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}

func (p tApplication) EnterDemo(
		user string,
		demo string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "demo", demo)
	return revel.MainRouter.Reverse("Application.EnterDemo", args).Url
}


type tWebSocket struct {}
var WebSocket tWebSocket


func (p tWebSocket) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("WebSocket.Room", args).Url
}

func (p tWebSocket) RoomSocket(
		user string,
		ws interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("WebSocket.RoomSocket", args).Url
}


type tLongPolling struct {}
var LongPolling tLongPolling


func (p tLongPolling) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("LongPolling.Room", args).Url
}

func (p tLongPolling) Say(
		user string,
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("LongPolling.Say", args).Url
}

func (p tLongPolling) WaitMessages(
		lastReceived int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "lastReceived", lastReceived)
	return revel.MainRouter.Reverse("LongPolling.WaitMessages", args).Url
}

func (p tLongPolling) Leave(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("LongPolling.Leave", args).Url
}


type tRefresh struct {}
var Refresh tRefresh


func (p tRefresh) Index(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Index", args).Url
}

func (p tRefresh) Room(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Room", args).Url
}

func (p tRefresh) Say(
		user string,
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("Refresh.Say", args).Url
}

func (p tRefresh) Leave(
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Refresh.Leave", args).Url
}


type tStatic struct {}
var Static tStatic


func (p tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (p tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


