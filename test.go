package albert

import "fmt"

func main() {
	i := 0

	go func() {
		i++ // write i
	}()

	fmt.Println(i) // read i
}
