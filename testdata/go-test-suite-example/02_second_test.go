package main

import (
	"testing"
)

/*
<speck>
### Test 2
</speck>
*/

func Test02Second(t *testing.T) {
	/*
	<speck tab=1>
	In this test we will ensure that the letter "a" is indeed equal to itself.
	</speck>
	*/
	if "a" != "a" {
		/*
		<speck tab=2>
		#### Possible Failures

		In rare cases, the letter "a" might equal something else, in which case
		you may want to question the entire existence of the universe.
		</speck>
		*/
		t.Errorf("How can \"a\" not equal \"a\"?")
	}
}
