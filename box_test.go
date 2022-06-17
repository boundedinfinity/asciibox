package asciibox_test

import (
	"testing"

	"github.com/boundedinfinity/asciibox"
	"github.com/boundedinfinity/asciibox/alignment_type"
	"github.com/stretchr/testify/assert"
)

var (
	input = []string{
		"a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z",
		"test message",
		"on mulitple",
		"lines",
	}
)

func Test_Box_align_left(t *testing.T) {
	expected := `//*******************
//*                 *
//* a b c d e f g h *
//* i j k l m n o   *
//* p q r s t u v w *
//* x y z a b c d   *
//* e f g h i j k l *
//* m n o p q r s   *
//* t u v w x y z   *
//* test message    *
//* on mulitple     *
//* lines           *
//*                 *
//*******************`

	actual := asciibox.Box(input, asciibox.BoxOptions{
		Alignment: alignment_type.Left,
		BoxWidth:  15,
	})

	assert.Equal(t, expected, actual)
}

func Test_Box_align_middle(t *testing.T) {
	expected := `//*******************
//*                 *
//* a b c d e f g h *
//*  i j k l m n o  *
//* p q r s t u v w *
//*  x y z a b c d  *
//* e f g h i j k l *
//*  m n o p q r s  *
//*  t u v w x y z  *
//*  test message   *
//*   on mulitple   *
//*      lines      *
//*                 *
//*******************`

	actual := asciibox.Box(input, asciibox.BoxOptions{
		Alignment: alignment_type.Middle,
		BoxWidth:  15,
	})

	assert.Equal(t, expected, actual)
}

func Test_Box_align_right(t *testing.T) {
	expected := `//*******************
//*                 *
//* a b c d e f g h *
//*   i j k l m n o *
//* p q r s t u v w *
//*   x y z a b c d *
//* e f g h i j k l *
//*   m n o p q r s *
//*   t u v w x y z *
//*    test message *
//*     on mulitple *
//*           lines *
//*                 *
//*******************`

	actual := asciibox.Box(input, asciibox.BoxOptions{
		Alignment: alignment_type.Right,
		BoxWidth:  15,
	})

	assert.Equal(t, expected, actual)
}
