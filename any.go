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
|  QUEST 09 : Any                       |           |
'---------------------------------------'----------*/
package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, v := range arr {
		f(v)
	}
}