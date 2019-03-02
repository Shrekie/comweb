package input

// Input channels Data in a specific way and
// performs destructuring to then envoy the process
// to a potrayer
type Input struct {
	Channel     Channel
	Destructure Destructure
	Envoy       Envoy
}

// Transmit the invocation
func (i *Input) Transmit() (Envoy, error) {
	args := i.Channel.process(i.Destructure)
	query, err := args.query()

	if err != nil {
		return i.Envoy, err
	}

	i.Envoy.set(query)
	return i.Envoy, err
}
