package sttabwriter

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func TabWriter() {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRodomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJonh\tSmith\t")
	w.Flush()
}
