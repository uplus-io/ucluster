/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import (
	log "github.com/uplus-io/ugo/logger"
)

type EventListener interface {
	OnTopologyChanged(event *NodeEvent)
}

type ClusterEventListener struct {
	warehouse *Warehouse
}

func NewClusterEventListener(warehouse *Warehouse) *ClusterEventListener {
	return &ClusterEventListener{warehouse: warehouse}
}

func (p *ClusterEventListener) OnTopologyChanged(event *NodeEvent) {
	switch Type {
	case NodeEventType_Join:
		JoinNode(event.Node.Ip, int(event.Node.Port))
	case NodeEventType_Leave:
		LeaveNode(event.Node.Ip, int(event.Node.Port))
	case NodeEventType_Update:
		log.Debugf("node[%s] event update", event.Node.String())
	}
}
