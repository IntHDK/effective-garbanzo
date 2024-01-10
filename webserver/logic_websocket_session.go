package webserver

import (
	"sync"

	"github.com/gorilla/websocket"
)

type websocketSessionManager struct {
	wsconnections           *sync.Map //key:string(session unique id) value:*websocket.Conn(actual connection)
	subscriptions           *sync.Map //key:string(group id) value:*sync.Map[string]interface{} (key:session unique id)
	wsconnectionsessiondata *sync.Map //key:string(session unique id) value:*sync.Map[string]string (key:parameter, value:value)
}

func NewWebsocketSessionManager() (res *websocketSessionManager) {
	res = &websocketSessionManager{
		wsconnections:           new(sync.Map),
		subscriptions:           new(sync.Map),
		wsconnectionsessiondata: new(sync.Map),
	}
	return
}

func (wsm *websocketSessionManager) Register(connid string, conn *websocket.Conn) (isfirstregistered bool) {
	_, ext := wsm.wsconnections.LoadOrStore(connid, conn)
	return !ext
}
func (wsm *websocketSessionManager) Unregister(connid string) {
	wsm.wsconnections.Delete(connid)
}

// 빈 subscription 생성
func (wsm *websocketSessionManager) MakeGroup(groupid string) (isfirstcreated bool) {
	_, ext := wsm.subscriptions.LoadOrStore(groupid, new(sync.Map))
	isfirstcreated = !ext
	return
}
func (wsm *websocketSessionManager) SubscribeGroup(groupid string, connid string) (ok bool) {
	groupsrc, _ := wsm.subscriptions.LoadOrStore(groupid, new(sync.Map))
	if groupset, oktype := groupsrc.(*sync.Map); oktype {
		groupset.LoadOrStore(connid, new(interface{}))
		return true
	}
	return false
}
func (wsm *websocketSessionManager) UnsubscribeGroup(groupid string, connid string) (ok bool) {
	groupsrc, _ := wsm.subscriptions.LoadOrStore(groupid, new(sync.Map))
	if groupset, oktype := groupsrc.(*sync.Map); oktype {
		groupset.LoadOrStore(connid, new(interface{}))
		return true
	}
	return false
}

// 비어있을때만 subscription 제거
func (wsm *websocketSessionManager) RemoveGroup(groupid string) (res bool) {
	src, ext := wsm.subscriptions.Load(groupid)
	if ext {
		if srcmap, oktype := src.(*sync.Map); oktype {
			cnt := 0
			srcmap.Range(func(k, v any) bool {
				cnt++
				return true
			})
			if cnt > 0 {
				return false
			} else {
				wsm.subscriptions.Delete(groupid)
			}
		} else {
			wsm.subscriptions.Delete(groupid)
			res = true
			return
		}
	}
	return
}
