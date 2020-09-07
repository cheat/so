package api

import (
	"testing"
)

// TestExtract asserts that Extract properly extracts code blocks from an
// Answer's body
func TestExtract(t *testing.T) {

	// initialize an answer
	// NB: important to note is that we only want to extract the
	// <pre><code></code></pre> blocks, not the inline <code></code> blocks.
	body := `
	<p>Try this:

	<pre><code>
	fmt.Println("hello, world!")
	</code></pre>

	<p>It will write <code>"hello, world!"</code> to stdout`
	ans := Answer{Body: body}

	// extract the code blocks
	got, err := ans.Extract()
	if err != nil {
		t.Errorf("failed to extract code: %v", err)
	}

	// assert that the expected result was returned
	want := `fmt.Println("hello, world!")`
	if got != want {
		t.Errorf("failed to extract code: want: %s, got: %s", want, got)
	}
}
