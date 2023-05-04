type handler interface {
	sendRequest(string) string
}

type handler1 struct {
	nextHandler handler
}

func (ha *handler1) sendRequest(msg string) string {
	switch {
	case msg == "H1":
		return fmt.Sprintf("Handler A done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return "err"
	}
}

type handler2 struct {
	nextHandler handler
}

func (ha *handler2) sendRequest(msg string) string {
	switch {
	case msg == "H2":
		return fmt.Sprintf("Handler B done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return "err"
	}
}

type handler3 struct {
	nextHandler handler
}

func (ha *handler3) sendRequest(msg string) string {
	switch {
	case msg == "H3":
		return fmt.Sprintf("Handler C done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return "err"
	}
}

func main() {
	handl1 := &handler1{}
	handl2 := &handler2{}
	handl3 := &handler3{}
	handl1.nextHandler = handl2
	handl2.nextHandler = handl3

	req := "H2"
	fmt.Println(handl1.sendRequest(req))
}
