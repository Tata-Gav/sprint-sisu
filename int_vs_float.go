package sprint

func IntVsFloat(i int, f float32) string {

	floatI := float32(i)

	if floatI > f {
		return "Integer"
	} else if floatI < f {
		return "Float"
	}else {
		return "Same"
	}
}