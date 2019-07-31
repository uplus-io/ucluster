module github.com/uplus-io/ucluster

go 1.12

require (
	github.com/hashicorp/go-sockaddr v1.0.0
	github.com/hashicorp/memberlist v0.1.4
	github.com/uplus-io/uengine v0.0.0
	github.com/uplus-io/ugo v0.0.0
)

replace (
	github.com/uplus-io/uengine => ../uengine
	github.com/uplus-io/ugo => ../ugo
)
