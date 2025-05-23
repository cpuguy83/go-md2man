package md2man

import (
	"testing"

	"github.com/russross/blackfriday/v2"
)

type TestParams struct {
	extensions blackfriday.Extensions
}

func TestCodeBlocks(t *testing.T) {
	tests := []string{
		"```\nsome code\n```\n",
		".nh\n\n.EX\nsome code\n.EE\n",

		"```bash\necho foo\n```\n",
		".nh\n\n.EX\necho foo\n.EE\n",

		// make sure literal new lines surrounding the markdown block are preserved as they are intentional
		"```bash\n\nsome code\n\n```",
		".nh\n\n.EX\n\nsome code\n\n.EE\n",
	}
	doTestsParam(t, tests, TestParams{blackfriday.FencedCode})
}

func TestEmphasis(t *testing.T) {
	tests := []string{
		"nothing inline\n",
		".nh\n\n.PP\nnothing inline\n",

		"simple *inline* test\n",
		".nh\n\n.PP\nsimple \\fIinline\\fP test\n",

		"*at the* beginning\n",
		".nh\n\n.PP\n\\fIat the\\fP beginning\n",

		"at the *end*\n",
		".nh\n\n.PP\nat the \\fIend\\fP\n",

		"*try two* in *one line*\n",
		".nh\n\n.PP\n\\fItry two\\fP in \\fIone line\\fP\n",

		"over *two\nlines* test\n",
		".nh\n\n.PP\nover \\fItwo\nlines\\fP test\n",

		"odd *number of* markers* here\n",
		".nh\n\n.PP\nodd \\fInumber of\\fP markers* here\n",

		"odd *number\nof* markers* here\n",
		".nh\n\n.PP\nodd \\fInumber\nof\\fP markers* here\n",

		"simple _inline_ test\n",
		".nh\n\n.PP\nsimple \\fIinline\\fP test\n",

		"_at the_ beginning\n",
		".nh\n\n.PP\n\\fIat the\\fP beginning\n",

		"at the _end_\n",
		".nh\n\n.PP\nat the \\fIend\\fP\n",

		"_try two_ in _one line_\n",
		".nh\n\n.PP\n\\fItry two\\fP in \\fIone line\\fP\n",

		"over _two\nlines_ test\n",
		".nh\n\n.PP\nover \\fItwo\nlines\\fP test\n",

		"odd _number of_ markers_ here\n",
		".nh\n\n.PP\nodd \\fInumber of\\fP markers_ here\n",

		"odd _number\nof_ markers_ here\n",
		".nh\n\n.PP\nodd \\fInumber\nof\\fP markers_ here\n",

		"mix of *markers_\n",
		".nh\n\n.PP\nmix of *markers_\n",

		"*What is A\\* algorithm?*\n",
		".nh\n\n.PP\n\\fIWhat is A* algorithm?\\fP\n",
	}
	doTestsInline(t, tests)
}

func TestStrong(t *testing.T) {
	tests := []string{
		"nothing inline\n",
		".nh\n\n.PP\nnothing inline\n",

		"simple **inline** test\n",
		".nh\n\n.PP\nsimple \\fBinline\\fP test\n",

		"**at the** beginning\n",
		".nh\n\n.PP\n\\fBat the\\fP beginning\n",

		"at the **end**\n",
		".nh\n\n.PP\nat the \\fBend\\fP\n",

		"**try two** in **one line**\n",
		".nh\n\n.PP\n\\fBtry two\\fP in \\fBone line\\fP\n",

		"over **two\nlines** test\n",
		".nh\n\n.PP\nover \\fBtwo\nlines\\fP test\n",

		"odd **number of** markers** here\n",
		".nh\n\n.PP\nodd \\fBnumber of\\fP markers** here\n",

		"odd **number\nof** markers** here\n",
		".nh\n\n.PP\nodd \\fBnumber\nof\\fP markers** here\n",

		"simple __inline__ test\n",
		".nh\n\n.PP\nsimple \\fBinline\\fP test\n",

		"__at the__ beginning\n",
		".nh\n\n.PP\n\\fBat the\\fP beginning\n",

		"at the __end__\n",
		".nh\n\n.PP\nat the \\fBend\\fP\n",

		"__try two__ in __one line__\n",
		".nh\n\n.PP\n\\fBtry two\\fP in \\fBone line\\fP\n",

		"over __two\nlines__ test\n",
		".nh\n\n.PP\nover \\fBtwo\nlines\\fP test\n",

		"odd __number of__ markers__ here\n",
		".nh\n\n.PP\nodd \\fBnumber of\\fP markers__ here\n",

		"odd __number\nof__ markers__ here\n",
		".nh\n\n.PP\nodd \\fBnumber\nof\\fP markers__ here\n",

		"mix of **markers__\n",
		".nh\n\n.PP\nmix of **markers__\n",

		"**`/usr`** : this folder is named `usr`\n",
		".nh\n\n.PP\n\\fB\\fB/usr\\fR\\fP : this folder is named \\fBusr\\fR\n",

		"**`/usr`** :\n\n this folder is named `usr`\n",
		".nh\n\n.PP\n\\fB\\fB/usr\\fR\\fP :\n\n.PP\nthis folder is named \\fBusr\\fR\n",
	}
	doTestsInline(t, tests)
}

func TestEmphasisMix(t *testing.T) {
	tests := []string{
		"***triple emphasis***\n",
		".nh\n\n.PP\n\\fB\\fItriple emphasis\\fP\\fP\n",

		"***triple\nemphasis***\n",
		".nh\n\n.PP\n\\fB\\fItriple\nemphasis\\fP\\fP\n",

		"___triple emphasis___\n",
		".nh\n\n.PP\n\\fB\\fItriple emphasis\\fP\\fP\n",

		"***triple emphasis___\n",
		".nh\n\n.PP\n***triple emphasis___\n",

		"*__triple emphasis__*\n",
		".nh\n\n.PP\n\\fI\\fBtriple emphasis\\fP\\fP\n",

		"__*triple emphasis*__\n",
		".nh\n\n.PP\n\\fB\\fItriple emphasis\\fP\\fP\n",

		"**improper *nesting** is* bad\n",
		".nh\n\n.PP\n\\fBimproper *nesting\\fP is* bad\n",

		"*improper **nesting* is** bad\n",
		".nh\n\n.PP\n*improper \\fBnesting* is\\fP bad\n",
	}
	doTestsInline(t, tests)
}

func TestCodeSpan(t *testing.T) {
	tests := []string{
		"`source code`\n",
		".nh\n\n.PP\n\\fBsource code\\fR\n",

		"` source code with spaces `\n",
		".nh\n\n.PP\n\\fBsource code with spaces\\fR\n",

		"` source code with spaces `not here\n",
		".nh\n\n.PP\n\\fBsource code with spaces\\fRnot here\n",

		"a `single marker\n",
		".nh\n\n.PP\na `single marker\n",

		"a single multi-tick marker with ``` no text\n",
		".nh\n\n.PP\na single multi-tick marker with ``` no text\n",

		"markers with ` ` a space\n",
		".nh\n\n.PP\nmarkers with  a space\n",

		"`source code` and a `stray\n",
		".nh\n\n.PP\n\\fBsource code\\fR and a `stray\n",

		"`source *with* _awkward characters_ in it`\n",
		".nh\n\n.PP\n\\fBsource *with* _awkward characters_ in it\\fR\n",

		"`split over\ntwo lines`\n",
		".nh\n\n.PP\n\\fBsplit over\ntwo lines\\fR\n",

		"```multiple ticks``` for the marker\n",
		".nh\n\n.PP\n\\fBmultiple ticks\\fR for the marker\n",

		"```multiple ticks `with` ticks inside```\n",
		".nh\n\n.PP\n\\fBmultiple ticks `with` ticks inside\\fR\n",
	}
	doTestsInline(t, tests)
}

func TestFlatLists(t *testing.T) {
	tests := []string{
		"Paragraph\n\n- item one\n- item two\n- item three\n",
		".nh\n\n.PP\nParagraph\n.IP \\(bu 2\nitem one\n.IP \\(bu 2\nitem two\n.IP \\(bu 2\nitem three\n",
	}
	doTestsInline(t, tests)
}

func TestListLists(t *testing.T) {
	tests := []string{
		"\n\n**[grpc]**\n: Section for gRPC socket listener settings. Contains three properties:\n - **address** (Default: \"/run/containerd/containerd.sock\")\n - **uid** (Default: 0)\n - **gid** (Default: 0)",
		".nh\n\n.TP\n\\fB[grpc]\\fP\nSection for gRPC socket listener settings. Contains three properties:\n.RS\n.IP \\(bu 2\n\\fBaddress\\fP (Default: \"/run/containerd/containerd.sock\")\n.IP \\(bu 2\n\\fBuid\\fP (Default: 0)\n.IP \\(bu 2\n\\fBgid\\fP (Default: 0)\n.RE\n",
		"Definition title\n: Definition description one\n: And two\n: And three\n",
		".nh\n\n.TP\nDefinition title\nDefinition description one\n\nAnd two\n\nAnd three\n",
		"Definition\n:    description\n\n     split\n\n     over\n\n     multiple\n\n     paragraphs\n",
		".nh\n\n.TP\nDefinition\ndescription\n\nsplit\n\nover\n\nmultiple\n\nparagraphs\n",
	}
	doTestsParam(t, tests, TestParams{blackfriday.DefinitionLists})
}

func TestLineBreak(t *testing.T) {
	tests := []string{
		"this line  \nhas a break\n",
		".nh\n\n.PP\nthis line\n.br\nhas a break\n",

		"this line \ndoes not\n",
		".nh\n\n.PP\nthis line\ndoes not\n",

		"this line\\\ndoes not\n",
		".nh\n\n.PP\nthis line\\\\\ndoes not\n",

		"this line\\ \ndoes not\n",
		".nh\n\n.PP\nthis line\\\\\ndoes not\n",

		"this has an   \nextra space\n",
		".nh\n\n.PP\nthis has an\n.br\nextra space\n",
	}
	doTestsInline(t, tests)

	tests = []string{
		"this line  \nhas a break\n",
		".nh\n\n.PP\nthis line\n.br\nhas a break\n",

		"this line \ndoes not\n",
		".nh\n\n.PP\nthis line\ndoes not\n",

		"this line\\\nhas a break\n",
		".nh\n\n.PP\nthis line\n.br\nhas a break\n",

		"this line\\ \ndoes not\n",
		".nh\n\n.PP\nthis line\\\\\ndoes not\n",

		"this has an   \nextra space\n",
		".nh\n\n.PP\nthis has an\n.br\nextra space\n",
	}
	doTestsInlineParam(t, tests, TestParams{
		extensions: blackfriday.BackslashLineBreak,
	})
}

func TestTable(t *testing.T) {
	tests := []string{
		`
| Animal               | Color         |
| --------------| --- |
| elephant        | Gray. The elephant is very gray.  |
| wombat     | No idea.      |
| zebra        | Sometimes black and sometimes white, depending on the stripe.     |
| robin | red. |
| Meerschweinchen a.k.a. guinea pig | Varies. |
`,
		`'\" t
.nh

.TS
allbox;
l l 
l l .
\fBAnimal\fP	\fBColor\fP
elephant	T{
Gray. The elephant is very gray.
T}
wombat	No idea.
zebra	T{
Sometimes black and sometimes white, depending on the stripe.
T}
robin	red.
T{
Meerschweinchen a.k.a. guinea pig
T}	Varies.
.TE
`,
	}
	doTestsInlineParam(t, tests, TestParams{blackfriday.Tables})
}

func TestTableWithEmptyCell(t *testing.T) {
	tests := []string{
		`
| Col1     | Col2  | Col3 |
|:---------|:-----:|:----:| 
| row one  |       |      | 
| row two  | x     |      |
`,
		`'\" t
.nh

.TS
allbox;
l l l 
l l l .
\fBCol1\fP	\fBCol2\fP	\fBCol3\fP
row one		
row two	x	
.TE
`,
	}
	doTestsInlineParam(t, tests, TestParams{blackfriday.Tables})
}

func TestTableWrapping(t *testing.T) {
	tests := []string{
		`
| Col1        | Col2                                             |
| ----------- | ------------------------------------------------ |
| row one     | This is a short line.                            |
| row\|two    | Col1 should not wrap.                            |
| row three   | no\|wrap                                         |
| row four    | Inline _cursive_ should not wrap.                |
| row five    | Inline ` + "`code markup`" + ` should not wrap.  |
| row six     | A line that's longer than 30 characters with inline ` + "`code markup`" + ` or _cursive_ should not wrap.  |
| row seven   | Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent eu ipsum eget tortor aliquam accumsan. Quisque ac turpis convallis, sagittis urna ac, tempor est. Mauris nibh arcu, hendrerit id eros sed, sodales lacinia ex. Suspendisse sed condimentum urna, vitae mattis lectus. Mauris imperdiet magna vel purus pretium, id interdum libero. |
`,
		`'\" t
.nh

.TS
allbox;
l l 
l l .
\fBCol1\fP	\fBCol2\fP
row one	This is a short line.
row|two	Col1 should not wrap.
row three	no|wrap
row four	Inline \fIcursive\fP should not wrap.
row five	Inline \fBcode markup\fR should not wrap.
row six	T{
A line that's longer than 30 characters with inline \fBcode markup\fR or \fIcursive\fP should not wrap.
T}
row seven	T{
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent eu ipsum eget tortor aliquam accumsan. Quisque ac turpis convallis, sagittis urna ac, tempor est. Mauris nibh arcu, hendrerit id eros sed, sodales lacinia ex. Suspendisse sed condimentum urna, vitae mattis lectus. Mauris imperdiet magna vel purus pretium, id interdum libero.
T}
.TE
`,
	}
	doTestsInlineParam(t, tests, TestParams{blackfriday.Tables})
}

func TestLinks(t *testing.T) {
	tests := []string{
		"See [docs](https://docs.docker.com/) for\nmore",
		".nh\n\n.PP\nSee docs\n\\[la]https://docs.docker.com/\\[ra] for\nmore\n",
		"See [docs](https://docs-foo.docker.com/) for\nmore",
		".nh\n\n.PP\nSee docs\n\\[la]https://docs\\-foo.docker.com/\\[ra] for\nmore\n",
		"See <https://docs-foo.docker.com/> for\nmore",
		".nh\n\n.PP\nSee \n\\[la]https://docs\\-foo.docker.com/\\[ra] for\nmore\n",
	}
	doTestsInline(t, tests)
}

func TestEscapeCharacters(t *testing.T) {
	tests := []string{
		"Test-one_two&three\\four~five",
		".nh\n\n.PP\nTest-one_two&three\\\\four~five\n",
		"'foo'\n'bar'",
		".nh\n\n.PP\n\\&'foo'\n\\&'bar'\n",
	}
	doTestsInline(t, tests)
}

func TestSpan(t *testing.T) {
	tests := []string{
		"Text containing a <span>html span</span> element\n",
		".nh\n\n.PP\nText containing a html span element\n",

		`Text containing an inline <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg"><image href="https://mdn.mozillademos.org/files/6457/mdn_logo_only_color.png" height="200" width="200"/></svg>SVG image`,
		".nh\n\n.PP\nText containing an inline SVG image\n",

		"Text containing a <span id=\"e-123\" class=\"foo\">html span</span> element\n",
		".nh\n\n.PP\nText containing a html span element\n",
	}
	doTestsInline(t, tests)
}

func TestEmails(t *testing.T) {
	tests := []string{
		`April 2014, Originally compiled by William Henry (whenry at redhat dot com)
based on docker.com source material and internal work.
June 2014, updated by Sven Dowideit <SvenDowideit@home.org.au>
July 2014, updated by Sven Dowideit (SvenDowideit@home.org.au)
`,
		`.nh

.PP
April 2014, Originally compiled by William Henry (whenry at redhat dot com)
based on docker.com source material and internal work.
June 2014, updated by Sven Dowideit SvenDowideit@home.org.au
\[la]mailto:SvenDowideit@home.org.au\[ra]
July 2014, updated by Sven Dowideit (SvenDowideit@home.org.au)
`,
	}
	doTestsInline(t, tests)
}

func TestComments(t *testing.T) {
	blockTests := []string{
		"First paragraph\n\n<!-- Comment, HTML should be separated by blank lines -->\n\nSecond paragraph\n",
		".nh\n\n.PP\nFirst paragraph\n\n.PP\nSecond paragraph\n",
	}
	doTestsParam(t, blockTests, TestParams{})

	inlineTests := []string{
		"Text with a com<!--...-->ment in the middle\n",
		".nh\n\n.PP\nText with a comment in the middle\n",
	}
	doTestsInlineParam(t, inlineTests, TestParams{})
}

func TestHeadings(t *testing.T) {
	tests := []string{
		"# title\n\n# NAME\ncommand - description\n\n# SYNOPSIS\nA short description\n\nWhich spans multiple paragraphs\n",
		".nh\n.TH title\n\n.SH NAME\ncommand \\- description\n\n\n.SH SYNOPSIS\nA short description\n\n.PP\nWhich spans multiple paragraphs\n",

		"# title\n\n# Name\nmy-command, other - description - with - hyphens\n",
		".nh\n.TH title\n\n.SH Name\nmy-command, other \\- description - with - hyphens\n",

		"# title\n\n# Not NAME\nsome - other - text\n",
		".nh\n.TH title\n\n.SH Not NAME\nsome - other - text\n",
	}
	doTestsInline(t, tests)
}

func execRecoverableTestSuite(t *testing.T, suite func(candidate *string)) {
	// Catch and report panics. This is useful when running 'go test -v' on
	// the integration server. When developing, though, crash dump is often
	// preferable, so recovery can be easily turned off with doRecover = false.
	var candidate string
	const doRecover = true
	if doRecover {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("\npanic while processing [%#v]: %s\n", candidate, err)
			}
		}()
	}
	suite(&candidate)
}

func runMarkdown(input string, params TestParams) string {
	renderer := NewRoffRenderer()
	return string(blackfriday.Run([]byte(input), blackfriday.WithRenderer(renderer),
		blackfriday.WithExtensions(params.extensions)))
}

func doTestsParam(t *testing.T, tests []string, params TestParams) {
	execRecoverableTestSuite(t, func(candidate *string) {
		for i := 0; i+1 < len(tests); i += 2 {
			input := tests[i]
			t.Run(input, func(t *testing.T) {
				*candidate = input
				expected := tests[i+1]
				actual := runMarkdown(*candidate, params)
				if actual != expected {
					t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
						*candidate, expected, actual)
				}

				// now test every substring to stress test bounds checking
				if !testing.Short() {
					for start := 0; start < len(input); start++ {
						for end := start + 1; end <= len(input); end++ {
							*candidate = input[start:end]
							runMarkdown(*candidate, params)
						}
					}
				}
			})
		}
	})
}

func doTestsInline(t *testing.T, tests []string) {
	doTestsInlineParam(t, tests, TestParams{})
}

func doTestsInlineParam(t *testing.T, tests []string, params TestParams) {
	params.extensions |= blackfriday.Strikethrough
	doTestsParam(t, tests, params)
}
