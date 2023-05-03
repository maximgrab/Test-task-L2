package main

var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}
	outchan := make(chan any)
	donechan := make(chan struct{})
	chanF := func(input <-chan interface{}) {
		for i := range input {
			outchan <- i
		}
		donechan <- struct{}{}
	}
	for _, i := range channels {
		go chanF(i)
	}
	go func() {
		<-donechan
		close(outchan)
	}()
	return outchan
}

func main() {

}
