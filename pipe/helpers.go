package pipe

import (
	"bytes"
	"github.com/fuzzy/spm/gout"
)

// string helper
func StrAppend(t gout.String, a gout.String) gout.String {
	var retv bytes.Buffer
	for i := 0; i < len(t); i++ {
		retv.WriteByte(t[i])
	}
	for i := 0; i < len(a); i++ {
		retv.WriteByte(a[i])
	}
	return gout.String(retv.String())
}
