package node

type Publisher struct {
	connector *Connector
}

func NewPublisher(conn *Connector) *Publisher {
	return &Publisher{
		connector: conn,
	}
}

func Publish(data interface{}) error {
	return nil
}
