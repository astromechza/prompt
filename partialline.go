/* All this function and file does is ensure that the cursor is moved down onto
its own line.

This is needed when the previous command's output did not end with a newline or left the cursor in a weird position.
*/

package main

import (
	"fmt"
)

func FixCursor(indicator string) {
	_, x, err := cursorPosition()
	if err != nil {
		panic(err)
	}
	if x > 0 {
		fmt.Println(indicator)
	}
}
