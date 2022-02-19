package gqltest

type Scenario struct {
	Name string
	Play func(*Env)
}
