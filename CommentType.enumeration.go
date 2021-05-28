////////////////////////////////////////////////////////////////////////
// Generated by github.com/boundedinfinity/enumer
////////////////////////////////////////////////////////////////////////

package asciibox

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CommentType string

const (
	CommentType_DoubleBackslash CommentType = "double_backslash"
	CommentType_Pound           CommentType = "pound"
)

var (
	CommentTypes = []CommentType{
		CommentType_DoubleBackslash,
		CommentType_Pound,
	}
)

func IsCommentType(v string) bool {
	var f bool

	for _, e := range CommentTypes {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func CommentTypeParse(v string) (CommentType, error) {
	var o CommentType
	var f bool
	n := strings.ToLower(v)

	for _, e := range CommentTypes {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, ErrCommentTypeNotFound(v)
	}

	return o, nil
}

func ErrCommentTypeNotFound(v string) error {
	var ss []string

	for _, e := range CommentTypes {
		ss = append(ss, string(e))
	}

	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(ss, ","),
	)
}

func (t CommentType) String() string {
	return string(t)
}

func (t CommentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *CommentType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := CommentTypeParse(s)

	if err != nil {
		return err
	}

	t = &e

	return nil
}