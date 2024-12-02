package p2p

// Peer is an interface that represents the remote node.
//
// 对等点是一个代表远程节点的接口。
type Peer interface{}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, websockets, ...)
//
// 是处理通信的任何东西
// 在网络中的节点之间。这可以是
// 形式（TCP、UDP、websockets 等）
type Transport interface {
	ListenAndAccept() error
}
