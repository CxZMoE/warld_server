package nova

import "fmt"

// A Object has a name, a position and abilities,thats all the define a thing.
type Object2D struct {
	Name      string
	Position  Vector2D
	Abilities []ObjectAbility
}

// Move the object to a new position
func (o *Object2D) MoveTo(v Vector2D) {
	o.Position = v
}

// Tag a object
func (o *Object2D) Tag(name string) {
	o.Name = name
}

// Abilities
type ObjectAbility struct {
	Name           string
	AbilityHandler func(args ...interface{}) // Called whenever a ability is triggered.
}

func (o *Object2D) AddAbility(ability ObjectAbility) {
	oa := new(ObjectAbility)
	oa.AbilityHandler = func(args ...interface{}) {
		fmt.Println(args[0].(ObjectAbility).Name)
	}
	o.Abilities = append(o.Abilities, ability)
}
func (o *Object2D) FindAbility(name string) (*ObjectAbility, int) {
	for i := range o.Abilities {
		if o.Abilities[i].Name == name {
			return &o.Abilities[i], i
		}
	}
	return nil, -1
}
func (o *Object2D) DelAbility(name string) {
	_, index := o.FindAbility(name)
	o.Abilities = append(o.Abilities[:index], o.Abilities[index+1:]...)
}
