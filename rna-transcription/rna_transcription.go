package strand

import "strings"

// ToRNA to convert dna to rna
func ToRNA(dna string) string {
	var rnaMap = map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	rna := ""
	dnas := strings.Split(dna, "")

	for _, d := range dnas {
		rna += rnaMap[d]
	}

	return rna
}
