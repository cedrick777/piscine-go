/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+																		+
+		**		**********	* ****** *	**********	   *****			+
+	**		**	**		**	*** ** ***	**	**	**	***		***			+
+	**		**	***				**		**		**	***		***			+
+	** ****	**		** ***		**		**	**		***		***			+
+	**		**			**		**		**		**	***		***			+
+	**		**	**********		**		**		**	   *****			+
+																		+
+ +++++++++++++++++++++++++++++++++++++++++++++	+						+
+					ZONE-01						+						+
+ +++++++++++++++++++++++++++++++++++++++++++++	+						+
+	NOM : TOURE 								+						+
+	PRENOMS : Ahmed Christian Cédrick			+						+
+	QUEST 06 : sortparams						+						+
+	GITHUB : github.com/cedrick777				+						+
+ +++++++++++++++++++++++++++++++++++++++++++++ +++++++++++++++++++++++*/

package piscine

func AppendRange(min, max int) []int {
	var tableau []int
	for i := min; i < max ; i++ {
		tableau = append(tableau,i)
	}
	return tableau
}