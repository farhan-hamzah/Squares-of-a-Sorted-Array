package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, j, compalier int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		A[i] = A[i]*A[i]
	}
	for i = 0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i] > A[j]{
				compalier = A[i]
				A[i] = A[j]
				A[j] = compalier
			}
		}
	}
	for i = 0; i < n; i++{
		fmt.Print(A[i], " ")
	}
}