package geeCache

import pb "github.com/catbugdemo/geeCache/geeCache"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key
type PeerPicker interface {
	// PickPeer picks a peer according to the key
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer
type PeerGetter interface {
	// Get gets the value for a key
	Get(in *pb.Request, out *pb.Response) error
}
