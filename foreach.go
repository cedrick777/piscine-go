/*--------------------------------------------------.
|     **      ******  ********  *******     ****    |
|  **    **  *.          **     **   **   **    **  |
|  **    **     ***.     **     ** ***    **    **  |
|  ********  *      *    **     **   **   **    **  |
|  **    **   ******     **     **   **     ****    |
|---------------------------------------.           |
|          ZONE-01 (cedrick777)         |           |
|---------------------------------------|           |
|  NOM : TOURE Ahmed Christian Cédrick  |           |
|  QUEST 09 : ForEach                   |           |
'---------------------------------------'----------*/
package piscine

import (
	"fmt"
)

func ForEach(f func(int), arr []int) {
	for _, v := range arr {
		f(v)
	}
	fmt.Println()
}