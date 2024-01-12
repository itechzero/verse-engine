package provider

type Provider interface {
	Enabled() bool
	Init() error
	Close() error
}

type RunProvider interface {
	Provider
	Run() error
	IsRunning() bool
}

type AbstractProvider struct {
	Provider
}

func (p *AbstractProvider) Enabled() bool {
	return p != nil
}

func (p *AbstractProvider) Init() error {
	return nil
}

func (p *AbstractProvider) Close() error {
	return nil
}

type AbstractRunProvider struct {
	RunProvider

	running bool
}

func (p *AbstractRunProvider) SetRunning(running bool) {
	if p != nil {
		p.running = running
	}
}

func (p *AbstractRunProvider) Enabled() bool {
	return p != nil
}

func (p *AbstractRunProvider) Init() error {
	return nil
}

func (p *AbstractRunProvider) Close() error {
	p.SetRunning(false)

	return nil
}

func (p *AbstractRunProvider) IsRunning() bool {
	return p.running
}
