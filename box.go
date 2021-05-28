package asciibox

import "strings"

var (
	defaultWrapCharacter    = "*"
	defaultCommentCharacter = CommentType_DoubleBackslash
	defaultAlignment        = Alignment_Left
	defaultBoxWidth         = 80
	defaultTopBottomPadding = 1
	defaultLeftRightPadding = 1
	defaultPaddingCharacter = " "
)

type BoxOptions struct {
	WrapCharacter    string
	PaddingCharacter string
	CommentCharater  CommentType
	BoxWidth         int
	TopBottomPadding int
	LeftRightPadding int
	Alignment        Alignment
}

type boxContext struct {
	ms          []string
	innerW      int
	longestText int
	options     BoxOptions
}

func Box(ms []string, options BoxOptions) string {
	var ms1 []string

	ctx := processContext(ms, options)
	ms1 = build(ctx)

	return strings.Join(ms1, "\n")
}

func build(ctx boxContext) []string {
	var ms1 []string

	wrapLine := strings.Repeat(ctx.options.WrapCharacter, ctx.options.BoxWidth)
	blankLine := strings.Repeat(ctx.options.PaddingCharacter, ctx.options.BoxWidth)

	ms1 = append(ms1, wrap(wrapLine, ctx.options.WrapCharacter, ctx))

	for i := 0; i < ctx.options.TopBottomPadding; i++ {
		ms1 = append(ms1, wrap(blankLine, ctx.options.PaddingCharacter, ctx))
	}

	for _, m := range ctx.ms {
		var s string

		switch ctx.options.Alignment {
		case Alignment_Left:
			s = s + m
			p := ctx.options.BoxWidth - len(s)
			s = s + strings.Repeat(ctx.options.PaddingCharacter, p)
		case Alignment_Middle:
			bm := ctx.options.BoxWidth / 2
			tm := len(m) / 2
			ts := bm - tm
			s = s + strings.Repeat(ctx.options.PaddingCharacter, ts)
			s = s + m
			p := ctx.options.BoxWidth - len(s)
			s = s + strings.Repeat(ctx.options.PaddingCharacter, p)
		case Alignment_Right:
			ts := ctx.options.BoxWidth - len(m)
			s = s + strings.Repeat(ctx.options.PaddingCharacter, ts)
			s = s + m
		}

		s = wrap(s, ctx.options.PaddingCharacter, ctx)
		ms1 = append(ms1, s)
	}

	for i := 0; i < ctx.options.TopBottomPadding; i++ {
		ms1 = append(ms1, wrap(blankLine, ctx.options.PaddingCharacter, ctx))
	}

	ms1 = append(ms1, wrap(wrapLine, ctx.options.WrapCharacter, ctx))
	return ms1
}
