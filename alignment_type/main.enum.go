//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package alignment_type

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/commons/slices"
	"github.com/boundedinfinity/commons/strings"
)

var (
	All = []Alignment{
		Left,
		Middle,
		Right,
	}
)

func pred(s string) func(Alignment) bool {
	return func(v Alignment) bool {
		return string(v) == s
	}
}

func (t Alignment) String() string {
	return string(t)
}

func Parse(v string) (Alignment, error) {
	f, ok := slices.FindFn(All, pred(v))

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slices.ContainsFn(All, func(v Alignment) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, strings.Join(All, ","),
	)
}

func (t Alignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *Alignment) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t Alignment) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *Alignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	if err := unmarshal(&s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}