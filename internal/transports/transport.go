package transport

type Transport interface {
	Notify(message string) error
}

func NewTransport() {

}
