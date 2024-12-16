package utils

type Vector struct {
	X int
	Y int
}

func (n Vector) Add(other Vector) Vector {
	n.X += other.X
	n.Y += other.Y
	return n
}

func (n Vector) Remove(other Vector) Vector {
	n.X -= other.X
	n.Y -= other.Y
	return n
}

func (n Vector) Multiply(multi int) Vector {
	n.X *= multi
	n.Y *= multi
	return n
}

func (n Vector) Divide(div int) Vector {
	n.X /= div
	n.Y /= div
	return n
}

func (n Vector) Modulo(other Vector) Vector {
	n.X %= other.X
	n.Y %= other.Y
	if n.X < 0 {
		n.X = other.X + n.X
	}
	if n.Y < 0 {
		n.Y = other.Y + n.Y
	}
	return n
}
