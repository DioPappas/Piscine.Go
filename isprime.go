package piscine

// for i := 2; i <= 9 ; i++{
// if nb % i == 0 {
// return false
//}
//}
// return true

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
