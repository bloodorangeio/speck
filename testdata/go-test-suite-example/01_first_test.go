package main

import (
	"testing"
)

/*
<speck>
### Test 1
</speck>
*/

func Test01First(t *testing.T) {
	/*
	<speck tab=1>
	In this test we will ensure that the number one (1) is indeed equal to itself.

	You would think this could never go wrong, but *oh* you would be surprised.

	Have you ever watched Sesame Street? Great show.

	```
	                       \WWW/
	                       /   \
	                      /wwwww\
	                    _|  o_o  |_
	       \WWWWWWW/   (_   / \   _)
	     _.'` o_o `'._   |  \_/  |
	    (_    (_)    _)  : ~~~~~ :
	      '.'-...-'.'     \_____/
	       (`'---'`)      [     ]
	jgs     `"""""`       `"""""`
	```
	</speck>
	 */
	if 1 != 1 {
		/*
		<speck tab=2>
		#### Possible Failures

		In rare cases, the number one (1) might equal something else, in which case
		you are encouraged to use a different programming language entirely.
		</speck>
		*/
		t.Errorf("How can 1 not equal 1?")
	}
}
