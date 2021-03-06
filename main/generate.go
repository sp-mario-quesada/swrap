package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var Package, TypeName, Type, File string

func init() {
	flag.StringVar(&Package, "p", "swrap", "package name")
	flag.StringVar(&TypeName, "n", "SWrap", "type name")
	flag.StringVar(&Type, "t", "byte", "type name")
	flag.StringVar(&File, "f", "swrap.go", "file name")
	flag.Parse()
}

func main() {
	// define params
	var param = struct {
		Package, TypeName, Type, Upper string
	}{
		Package:  Package,
		TypeName: TypeName,
		Type:     Type,
		Upper:    strings.Title(Type),
	}

	// open output file
	fd, err := os.Create(File)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	// execute template
	t := template.Must(template.New("swrap").Parse(str))
	err = t.Execute(fd, param)
	if err != nil {
		log.Fatal(err)
	}
}

var str = `package {{.Package}}

type {{.TypeName}} []{{.Type}}

/**
 * Unshift()v       vPush(), Add()
 *         [ 1, 2, 3 ]
 *    Shift()^     ^Pop()
 */

// create swrap from []{{.Type}} and return
func New(a []{{.Type}}) {{.TypeName}} {
	return {{.TypeName}}(a)
}

// create swrap from []{{.Type}} and return address
func Make(a []{{.Type}}) *{{.TypeName}} {
	sw := {{.TypeName}}(a)
	return &sw
}

// get []{{.Type}} from swrap
func (sw *{{.TypeName}}) {{.Upper}}s() []{{.Type}} {
	return []{{.Type}}(*sw)
}

// get length of swrap
func (sw *{{.TypeName}}) Len() int {
	return len(*sw)
}

// add {{.Type}} at the end of swrap
func (sw *{{.TypeName}}) Add(a {{.Type}}) {
	*sw = append(*sw, a)
}

// merge given []{{.Type}} to swrap ([swrap.., []{{.Type}}...])
func (sw *{{.TypeName}}) Merge(a []{{.Type}}) {
	s := *sw
	l := len(s) + len(a)
	ss := make([]{{.Type}}, l, l)
	copy(ss[0:], s[:])
	copy(ss[len(s):], a)
	*sw = ss
}

// delete given index value from swrap
func (sw *{{.TypeName}}) Delete(i int) {
	s := *sw
	copy(s[i:], s[i+1:])
	// s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
}

// compare given []{{.Type}} with swrap
func (sw *{{.TypeName}}) Compare(b []{{.Type}}) bool {
	s := *sw
	if len(s) != len(b) {
		return false
	}

	for i, v := range b {
		if s[i] != v {
			return false
		}
	}
	return true
}

// add {{.Type}} at the end of swrap (alias of Add())
func (sw *{{.TypeName}}) Push(b {{.Type}}) {
	*sw = append(*sw, b)
}

// get {{.Type}} at the end of swrap
func (sw *{{.TypeName}}) Pop() {{.Type}} {
	s := *sw
	last := s[len(s)-1]
	s[len(s)-1] = 0 // GC
	*sw = s[:len(s)-1]
	return last
}

// add {{.Type}} at the top of swrap
func (sw *{{.TypeName}}) UnShift(b {{.Type}}) {
	s := *sw
	l := len(s) + 1
	ss := make([]{{.Type}}, l, l)
	ss[0] = b
	copy(ss[1:], s[:])
	*sw = ss
}

// get {{.Type}} at the top of swrap
func (sw *{{.TypeName}}) Shift() {{.Type}} {
	s := *sw
	top := s[0]
	s[0] = 0 // GC
	*sw = s[1:]
	return top
}

// replace given index value with given {{.Type}}
func (sw *{{.TypeName}}) Replace(i int, b {{.Type}}) {
	s := *sw
	over := i - len(s)
	if over > -1 {
		ss := make([]{{.Type}}, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = b
	*sw = s
}
`
