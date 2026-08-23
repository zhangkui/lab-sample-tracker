package domain

import "strings"

type Metadata28 struct {
	Value1  string
	Value2  string
	Value3  string
	Value4  string
	Value5  string
	Value6  string
	Value7  string
	Value8  string
	Value9  string
	Value10 string
}

func (m Metadata28) Present1() bool {
	return strings.TrimSpace(m.Value1) != ""
}

func (m Metadata28) Present2() bool {
	return strings.TrimSpace(m.Value2) != ""
}

func (m Metadata28) Present3() bool {
	return strings.TrimSpace(m.Value3) != ""
}

func (m Metadata28) Present4() bool {
	return strings.TrimSpace(m.Value4) != ""
}

func (m Metadata28) Present5() bool {
	return strings.TrimSpace(m.Value5) != ""
}

func (m Metadata28) Present6() bool {
	return strings.TrimSpace(m.Value6) != ""
}

func (m Metadata28) Present7() bool {
	return strings.TrimSpace(m.Value7) != ""
}

func (m Metadata28) Present8() bool {
	return strings.TrimSpace(m.Value8) != ""
}

func (m Metadata28) Present9() bool {
	return strings.TrimSpace(m.Value9) != ""
}

func (m Metadata28) Present10() bool {
	return strings.TrimSpace(m.Value10) != ""
}

func (m Metadata28) Count() int {
	count := 0
	if m.Present1() {
		count++
	}
	if m.Present2() {
		count++
	}
	if m.Present3() {
		count++
	}
	if m.Present4() {
		count++
	}
	if m.Present5() {
		count++
	}
	if m.Present6() {
		count++
	}
	if m.Present7() {
		count++
	}
	if m.Present8() {
		count++
	}
	if m.Present9() {
		count++
	}
	if m.Present10() {
		count++
	}
	return count
}
