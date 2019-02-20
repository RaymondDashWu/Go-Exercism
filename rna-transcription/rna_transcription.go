package strand

// Notes
// G -> C
// C -> G
// T -> A
// A -> U

func ToRNA(dna string) string {
	var emptyString = ""
	if len(dna) >= 1 {
		for i := 0; i < len(dna); i++ {
			if dna[i] == 'G' {
				emptyString += string('C')
			}
			if dna[i] == 'C' {
				emptyString += string('G')
			}
			if dna[i] == 'T' {
				emptyString += string('A')
			}
			if dna[i] == 'A' {
				emptyString += string('U')
			}
		}
	}
	return emptyString
}
