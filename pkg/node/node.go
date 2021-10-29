package node

type Node interface {
	//start the node
	Start() error
	//stop  the node
	Stop() error
	//get node name string
	Name() string
	//publish a message to topic
	Publish(interface{}) error
	//subscribe a topic with call back
	Subscribe(string, func([]byte)) error
}
