package einx

import (
	"github.com/Cyinx/einx/agent"
	"github.com/Cyinx/einx/component"
	"github.com/Cyinx/einx/event"
	"github.com/Cyinx/einx/lua"
	"github.com/Cyinx/einx/module"
	"github.com/Cyinx/einx/network"
	"github.com/Cyinx/einx/timer"
	"sync"
)

type EventMsg = event.EventMsg
type EventType = event.EventType
type Agent = agent.Agent
type AgentID = agent.AgentID
type Component = component.Component
type ComponentID = component.ComponentID
type ComponentMgr = component.ComponentMgr
type MsgHandler = module.MsgHandler
type RpcHandler = module.RpcHandler
type Module = module.Module
type SessionEventMsg = event.SessionEventMsg
type LuaRuntime = lua_state.LuaRuntime
type NetLinker = network.NetLinker
type ProtoTypeID = network.ProtoTypeID
type SessionMgr = network.SessionMgr
type SessionHandler = network.SessionHandler
type ITcpClientMgr = network.ITcpClientMgr
type ITcpServerMgr = network.ITcpServerMgr
type TimerHandler = timer.TimerHandler
type EventReceiver = event.EventReceiver

type einx struct {
	end_wait   sync.WaitGroup
	close_chan chan bool
}

func (this *einx) do_close() {
	module.Close()
}

func (this *einx) close() {
	this.end_wait.Wait()
}
