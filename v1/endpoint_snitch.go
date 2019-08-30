/*
 * Copyright (c) 2019 uplus.io
 */

package v1

type EndpointSnitch interface {
	DataCenter() uint32

	Rack() uint32
}
