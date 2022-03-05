package driver

import "gides/src/hashring"

type Cluster struct {
	HashSlot [64]*hashring.ServerNode
}
