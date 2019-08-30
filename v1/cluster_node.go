/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
)

type NodeEventType uint8

const (
	NodeEventType_Join NodeEventType = iota
	NodeEventType_Leave
	NodeEventType_Update
)

type NodeEvent struct {
	Type   NodeEventType
	Node   *model.Node
	Native interface{}
}

func NewNodeEvent(t NodeEventType, node *model.Node, native interface{}) *NodeEvent {
	return &NodeEvent{Type: t, Node: node, Native: native}
}
