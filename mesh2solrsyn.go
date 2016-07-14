/* See LICENSE file for copyright and license details. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shugyousha/mesh"
)

type prefices []string

func (p prefices) Len() int {
	return len(p)
}

func (p prefices) Less(i, j int) bool {
	return len(p[i]) > len(p[j])
}

func (p prefices) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var (
	meshrecs mesh.MeSHRecordsMap
)

func gethyposyns(path string, mn *mesh.MeSHNode) []string {
	var res []string
	suffs := mn.GetSamePrefix(path)

	for _, suf := range suffs {
		rec, ok := meshrecs[suf]
		if !ok {
			fmt.Fprintf(os.Stderr, "Suf: %q does not have a record! Please check.\n", suf)
		}
		for s, _ := range rec.Entries {
			res = append(res, s)
		}

		res = append(res, rec.MH)
	}
	return res
}

func main() {
	var mrslice []*mesh.MeSHRecord

	mn := mesh.NewNode(make(map[string]*mesh.MeSHNode, 5))
	meshrecs = make(mesh.MeSHRecordsMap, 30000)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Need a file to read the MeSH terms from as an argument. Exiting.\n")
		os.Exit(1)
	}

	r := bufio.NewReader(os.Stdin)
	mesh.ParseMeSHTree(r, *mn)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file: %s\n", err)
	}

	fr := bufio.NewReader(f)
	_, meshrecs = mesh.ParseToSliceAndMap(*fr, meshrecs, mrslice)

	//suff := mn.GetSamePrefix("G11.561.600.810.964.186.624"})
	for r, _ := range meshrecs {
		syns := gethyposyns(r, mn)
		rec, ok := meshrecs[r]
		if !ok {
			fmt.Fprintf(os.Stderr, "Record: %q does not exist! Please check.\n", r)
			os.Exit(1)
		}
		for s, _ := range rec.Entries {
			syns = append(syns, s)
		}

		if len(syns) > 0 {
			//fmt.Fprintf(os.Stderr, "Rec: %q %q, syns: %q\n", r, rec.MH, strings.Join(syns, ","))
			fmt.Printf("%s => %s\n", rec.MH, strings.Join(syns, ", "))
		}
	}
}
