package DNAtoRNA

import s "strings"

func DNAtoRNA(dna string) string {
	return s.Replace(dna, "T", "U", -1)
}

// Using a new replacer
// r := strings.NewReplacer("T", "U")
// rna := r.Replace(dna)
