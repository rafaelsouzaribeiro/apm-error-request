package mongodb

type Configs struct {
	Name string
}

func NewConfigs(name string) *Configs {

	return &Configs{
		Name: name,
	}
}
