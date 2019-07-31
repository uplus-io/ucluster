/*
 * Copyright (c) 2019 uplus.io
 */

package ucluster

type EndpointSnitch interface {
	DataCenter() uint32

	Rack() uint32
}
