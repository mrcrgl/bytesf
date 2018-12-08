package bytesf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTexts = []string{
	"Für Designer, Schriftsetzer, Layouter, Grafikenthusiasten und alle anderen.",
	"Ein Projekt startet und doch es gibt noch keinen Text, allerdings sollte das Layout schon bald präsentiert werden ... was tun?",
	"Hier können verschieden Varianten von Lorem ipsum Text heruntergeladen werden. Jedes Beispiel ist als reines Text- oder Worddokument (in .zip Format) verfügbar.",
	"Damit das Projekt gleich starten kann benutze einfach etwas Lorem ipsum - Blind-, Füll-, Dummy-, Nachahmungs-, Platzhaltertext.",
	"Generiere einfach soviel Lorem Ipsum Text wie du brauchst, kopiere und füge ihn in dein Layout als vorübergehenden Platzhalter ein. Somit sieht das Projekt ein Stückchen vollständiger aus als zuvor. Viel Spaß dabei.",
	"PC - per Klick auf die rechte Maustaste und dann speichern.",
}

func TestBufferPool(t *testing.T) {
	b := NewBufferPool(64, 256)

	for i, text := range testTexts {
		i := i
		text := text
		t.Run(fmt.Sprintf("Turn %d", i), func(t *testing.T) {
			t.Parallel()

			for n := 0; n <= 1000; n++ {
				buf := b.GetBuffer()

				fmt.Fprint(buf, text)

				assert.Equal(t, text, buf.String())

				b.PutBuffer(buf)
			}

		})
	}
}

func Benchmark_NotPooled(b *testing.B) {
	for i, text := range testTexts {
		b.Run(fmt.Sprintf("Text %d (%d bytes)", i, len(text)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				buf := new(bytes.Buffer)
				buf.Grow(128)
				buf.WriteString(text)
			}
		})
	}
}
func Benchmark_Pooled(b *testing.B) {
	for i, text := range testTexts {
		bf := NewBufferPool(128, 256)

		b.Run(fmt.Sprintf("Text %d (%d bytes)", i, len(text)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				buf := bf.GetBuffer()
				buf.WriteString(text)
				bf.PutBuffer(buf)
			}
		})
	}
}
