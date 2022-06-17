package asciibox

import (
	"strings"

	"github.com/boundedinfinity/asciibox/alignment_type"
	"github.com/boundedinfinity/asciibox/comment_type"
	"github.com/boundedinfinity/commons/slices"
)

var (
	commentTypeMap = map[comment_type.CommentType]string{
		comment_type.DoubleBackslash: "//",
		comment_type.Pound:           "#",
	}
)

func processContext(ms []string, options BoxOptions) boxContext {
	coptions := BoxOptions{
		CommentCharater:  orElseCommentType(options.CommentCharater, defaultCommentCharacter),
		WrapCharacter:    orElseS(options.WrapCharacter, defaultWrapCharacter),
		PaddingCharacter: orElseS(options.PaddingCharacter, defaultPaddingCharacter),
		Alignment:        orElseAlignment(options.Alignment, defaultAlignment),
		BoxWidth:         orElseI(options.BoxWidth, defaultBoxWidth),
		TopBottomPadding: orElseI(options.TopBottomPadding, defaultTopBottomPadding),
		LeftRightPadding: orElseI(options.LeftRightPadding, defaultLeftRightPadding),
	}

	var cms []string

	for _, m := range ms {
		ss := splitAfterNCharacters(m, coptions.BoxWidth)
		ss = slices.Map(ss, func(v string) string { return strings.Trim(v, " ") })
		cms = append(cms, ss...)
	}

	var mLens []int

	for _, m := range ms {
		mLens = append(mLens, len(m))
	}

	longestText := maxI(mLens...)

	ctx := boxContext{
		ms:          cms,
		longestText: longestText,
		innerW:      maxI(longestText, coptions.BoxWidth),
		options:     coptions,
	}

	return ctx
}

func splitAfterNCharacters(s string, n int) []string {
	var ss []string
	sLen := len(s)

	if sLen < n {
		ss = append(ss, s)
		return ss
	}

	var t string

	for i := 0; i < sLen; i++ {
		if t != "" && len(t)%n == 0 {
			ss = append(ss, t)
			t = ""
		}

		t = t + string(s[i])
	}

	if len(t) > 0 {
		ss = append(ss, t)
	}

	return ss
}

func wrap(m, p string, ctx boxContext) string {
	var o string

	o = o + commentTypeMap[ctx.options.CommentCharater]
	o = o + ctx.options.WrapCharacter
	o = o + p
	o = o + m
	o = o + p
	o = o + ctx.options.WrapCharacter

	return o
}

func prefix(ctx boxContext, withPadd bool) string {
	var b strings.Builder

	b.WriteString(commentTypeMap[ctx.options.CommentCharater])
	b.WriteString(ctx.options.WrapCharacter)

	if withPadd {
		b.WriteString(strings.Repeat(ctx.options.PaddingCharacter, ctx.options.LeftRightPadding))
	}

	return b.String()
}

func suffix(ctx boxContext, withPad bool) string {
	var b strings.Builder

	if withPad {
		b.WriteString(strings.Repeat(ctx.options.PaddingCharacter, ctx.options.LeftRightPadding))
	}

	b.WriteString(ctx.options.WrapCharacter)

	return b.String()
}

func orElseAlignment(v, d alignment_type.Alignment) alignment_type.Alignment {
	if string(v) != "" {
		return v
	} else {
		return d
	}
}

func orElseCommentType(v, d comment_type.CommentType) comment_type.CommentType {
	if string(v) != "" {
		return v
	} else {
		return d
	}
}

func orElseS(v, d string) string {
	if v != "" {
		return v
	} else {
		return d
	}
}

func orElseI(v, d int) int {
	if v != 0 {
		return v
	} else {
		return d
	}
}

func maxI(is ...int) int {
	var m int

	for _, i := range is {
		if i > m {
			m = i
		}
	}

	return m
}
