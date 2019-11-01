package enum

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestParser_GetEnum(t *testing.T) {
	cases := []struct {
		description      string
		fileData         []byte
		trimPrefix       string
		formatFunc       func(s string) string
		expectedEnums    []Enum
		ExpectedTypeName string
	}{
		{
			description: "an integer enum",
			fileData: []byte(`package main

				//go:generate go-enum --trim "Post" --format snake
				const (
					PostCreate PostType = 0
					PostRead   PostType = 2
					PostUpdate PostType = 4
					PostDelete PostType = 8
				)`),
			trimPrefix: "Post",
			expectedEnums: []Enum{
				{
					Int:    0,
					Name:   "PostCreate",
					String: "CREATE",
				},
				{
					Int:    2,
					Name:   "PostRead",
					String: "READ",
				},
				{
					Int:    4,
					Name:   "PostUpdate",
					String: "UPDATE",
				},
				{
					Int:    8,
					Name:   "PostDelete",
					String: "DELETE",
				},
			},
			ExpectedTypeName: "PostType",
		},
		{
			description: "iota enum",
			fileData: []byte(`package main

				//go:generate go-enum --trim "Direction" --format upper
				const (
					DirectionUp DirectionType = iota
					DirectionDown
					DirectionLeft
					DirectionRight
				)`),
			trimPrefix: "Direction",
			expectedEnums: []Enum{
				{
					Int:    0,
					Name:   "DirectionUp",
					String: "UP",
				},
				{
					Int:    1,
					Name:   "DirectionDown",
					String: "DOWN",
				},
				{
					Int:    2,
					Name:   "DirectionLeft",
					String: "LEFT",
				},
				{
					Int:    3,
					Name:   "DirectionRight",
					String: "RIGHT",
				},
			},
			ExpectedTypeName: "DirectionType",
		},
		{
			description: "string enum",
			fileData: []byte(`package main

			//go:generate go-enum --trim "Answer" --format upper
			const (
				AnswerYes   YesOrNo = "Y"
				AnswerNo    YesOrNo = "N"
				AnswerMaybe YesOrNo = "M"
			)`),
			trimPrefix: "Answer",
			expectedEnums: []Enum{
				{
					Value:  "Y",
					Name:   "AnswerYes",
					String: "YES",
				},
				{
					Value:  "N",
					Name:   "AnswerNo",
					String: "NO",
				},
				{
					Value:  "M",
					Name:   "AnswerMaybe",
					String: "MAYBE",
				},
			},
			ExpectedTypeName: "YesOrNo",
		},
	}

	var (
		noFile  = ""
		noJSON  = false
		noValue = false
		lineNo  = 4
	)

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			parser := New(noFile, tc.trimPrefix, lineNo, noJSON, noValue, strings.ToUpper)
			err := parser.GetEnum(tc.fileData)

			require.NoError(t, err)

			assert.Equal(t, tc.ExpectedTypeName, parser.TypeName)
			assert.Equal(t, tc.expectedEnums, parser.Enums)
		})
	}
}
