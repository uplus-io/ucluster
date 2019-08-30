/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import "errors"

//cluster

var (
	ErrClusterNodeOffline = errors.New("cluster node offline")
)

var (
	ErrNotFoundClusterNode = errors.New("not found cluster node")

	ErrMessageDispatcherExist = errors.New("message dispatcher already exist")
	ErrMessageHandlerExist    = errors.New("message handler already exist")
	ErrMessageHandlerNotExist = errors.New("message handler not exist")
)
