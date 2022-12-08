package multiplikation

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	if x < 0 {
		return -Mult(-x, y)
	}
	if y < 0 {
		return -Mult(x, -y)
	}
	if x == 0 || y == 0 {
		return 0
	}

	return y + Mult(x-1, y)
	// return y + Mult(x,y) - y
}

/* Neben-Überlegungen

x * y == y + ((x-1) * y)
      == y + (x * y -y)

*/
