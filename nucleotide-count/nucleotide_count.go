package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[string]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = make(Histogram)
	for _, v := range d {
		// Possibly syntax issue? Logic for histogram is correct
		if v == h[string(v)] {
			h[string(v)]++
		}
		h[string(v)] = 1
	}
	// for i := 0; i < len(d); i++ {
	// 	if d[i] in h {

	// 	}
	// }

	return h, nil
}
