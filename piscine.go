package piscine

// UltimateDivMod ...
func UltimateDivMod(a *int, b *int) {
	*a = *a / *b
	*b = *a % *b
	// *a,*b = *a % / *b
}
