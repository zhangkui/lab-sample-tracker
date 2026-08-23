package domain

import "strings"

type AttributeSet24 struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 string
	Field5 string
	Field6 string
	Field7 string
	Field8 string
}

func (a AttributeSet24) Has1() bool {
	return strings.TrimSpace(a.Field1) != ""
}

func (a AttributeSet24) Has2() bool {
	return strings.TrimSpace(a.Field2) != ""
}

func (a AttributeSet24) Has3() bool {
	return strings.TrimSpace(a.Field3) != ""
}

func (a AttributeSet24) Has4() bool {
	return strings.TrimSpace(a.Field4) != ""
}

func (a AttributeSet24) Has5() bool {
	return strings.TrimSpace(a.Field5) != ""
}

func (a AttributeSet24) Has6() bool {
	return strings.TrimSpace(a.Field6) != ""
}

func (a AttributeSet24) Has7() bool {
	return strings.TrimSpace(a.Field7) != ""
}

func (a AttributeSet24) Has8() bool {
	return strings.TrimSpace(a.Field8) != ""
}

func (a AttributeSet24) Complete() bool {
	return a.Has1() && a.Has2() && a.Has3() && a.Has4() && a.Has5() && a.Has6() && a.Has7() && a.Has8()
}
