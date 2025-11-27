package connector

type Connector interface {
	Get() *Connection
	Release(*Connection)
}

type Connection struct {
	ID   int
	name string
}
