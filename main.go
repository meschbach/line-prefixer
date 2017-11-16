package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
		if( len( os.Args ) != 2 ){
			fmt.Println("Usage: " + os.Args[0] + " <line-token>")
			os.Exit( -1 )
		}

		var prefix = os.Args[1] + ": "
		var hasMore = true
    reader := bufio.NewReader(os.Stdin)

		for hasMore {
			text, err := reader.ReadString('\n')
			fmt.Print( prefix )
			fmt.Print( text )
			hasMore = err == nil
		}
		fmt.Print("\n")
}

