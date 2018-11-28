package merge

import "github.com/imdario/mergo"

// Merger merges two object into one overwriting
type Merger interface {
	Merge(dst, src interface{}) error
}

// MergoMerger mergo implementation of the merger
type MergoMerger struct{}

// Merge merges two object into one overwriting
func (m *MergoMerger) Merge(dst, src interface{}) error {
	return mergo.MergeWithOverwrite(dst, src)
}
