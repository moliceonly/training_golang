package basepractice


func ifelsecontrol(a int) int{
	b := 3
	if a < b {
		return a
	}else if a==4 {
		return a
	}else {
		return b
	}
}

func switchcontrol(a string)string{
	switch a{
	case "a":
		a = a + "a"
		fallthrough
	case "aa":
		a = a + "b"
	default:
		a=""
	}
	return a
}
