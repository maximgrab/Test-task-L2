type Button struct {
	command Command
}

type Device interface {
	On()
	Off()
}

type Printer struct {
	status bool
}

type OnCommand struct {
	device Device
}

type OffCommand struct {
	device Device
}

type Command interface {
	Execute()
}

func (pr *Printer) On() {
	pr.status = true
	fmt.Println("Printer on")
}

func (pr *Printer) Off() {
	pr.status = false
	fmt.Println("Printer off")
}

func (b *Button) Exec() {
	b.command.Execute()
}

func (on *OnCommand) Execute() {
	on.device.On()
}

func (off *OffCommand) Execute() {
	off.device.Off()
}

func main() {
	pr := &Printer{}
	on := &OnCommand{
		device: pr,
	}
	off := &OffCommand{
		device: pr,
	}
	buttonOn := &Button{
		command: on,
	}
	buttonOff := &Button{
		command: off,
	}
	buttonOn.Exec()
	buttonOff.Exec()
}
