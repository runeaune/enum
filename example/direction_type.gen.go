// Code generated by go-enum; DO NOT EDIT.
// Laste generated at 2019-11-12T09:46:37+01:00.
// go-enum: https://github.com/bombsimon/enum

package example

import (
	"encoding/json"
	"fmt"
)

// DirectionType is an Enum.
type DirectionType int

// DirectionTypeFromString returns a DirectionType from it's string representation.
func DirectionTypeFromString(s string) (DirectionType, error) {
	switch s {
	case DirectionUp.String():
		return DirectionUp, nil
	case DirectionDown.String():
		return DirectionDown, nil
	case DirectionLeft.String():
		return DirectionLeft, nil
	case DirectionRight.String():
		return DirectionRight, nil
	default:
		return DirectionType(-1), fmt.Errorf("unknown DirectionType %s", s)
	}
}

func (v DirectionType) String() string {
	switch v {
	case DirectionUp:
		return "UP"
	case DirectionDown:
		return "DOWN"
	case DirectionLeft:
		return "LEFT"
	case DirectionRight:
		return "RIGHT"
	default:
		return ""
	}
}

// Valid returns false if the DirectionType isn't valid.
func (v DirectionType) Valid() bool {
	switch v {
	case DirectionUp:
		return true
	case DirectionDown:
		return true
	case DirectionLeft:
		return true
	case DirectionRight:
		return true
	default:
		return false
	}
}

// MarshalJSON marshalls the DirectionType enum to it's JSON representation.
func (v DirectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// UnmarshalJSON unmarshalls the JSON to it's DirectionType enum.
func (v *DirectionType) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	newV, err := DirectionTypeFromString(s)
	if err != nil {
		return err
	}

	*v = newV

	return nil
}
