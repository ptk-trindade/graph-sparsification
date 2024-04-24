package effectiveresisitance

import (
	// "github.com/ptk-trindade/graph-sparsification/effectiveResistance/utils"
	"fmt"

	"github.com/ptk-trindade/graph-sparsification/utils"
	"gonum.org/v1/gonum/mat"
)

/*
https://www.universiteitleiden.nl/binaries/content/assets/science/mi/scripties/master/vos_vaya_master.pdf
Lv = c -> L is the Laplacian matrix, v is the voltage vector and c is the current vector
*/
func EffectiveResistance(adjList [][]int) *utils.EdgeWeight {
	edgeWeight := utils.NewEdgeWeight()
	laplacianMatrix := createLaplaceMatrix(adjList)

	// invert the Laplacian matrix
	var invLaplacian mat.Dense
	err := invLaplacian.Inverse(laplacianMatrix)
	if err != nil {
		fmt.Println("Error inverting Laplacian matrix")
		fmt.Println(err)
		return nil
	}

	// for each edge
	for i := 0; i < len(adjList); i++ {
		for j := 0; j < len(adjList[i]); j++ {
			nodeA := i
			nodeB := adjList[i][j]

			fmt.Println("a-b:", nodeA, nodeB)
			if nodeA > nodeB { // avoid computing edges twice
				continue
			}

			c_vals := make([]float64, len(adjList))
			c_vals[nodeA] = 1
			c_vals[nodeB] = -1

			c := mat.NewVecDense(len(c_vals), c_vals)
			fmt.Println("c:", c_vals)

			var v mat.VecDense
			v.MulVec(&invLaplacian, c)

			effecRes := v.At(nodeA, 0) - v.At(nodeB, 0)
			edgeWeight.AddWeight(nodeA, nodeB, effecRes)
		}
	}

	return edgeWeight
}

// This isn't exactly the Laplacian matrix, we add 1 to the first element of the matrix to make it non-singular
func createLaplaceMatrix(adjList [][]int) mat.Matrix {
	values := make([]float64, len(adjList)*len(adjList))

	for i := 0; i < len(adjList); i++ {
		for j := 0; j < len(adjList[i]); j++ {
			values[i*len(adjList)+adjList[i][j]] = -1
		}
		values[i*len(adjList)+i] = float64(len(adjList[i]))
	}

	values[0] += 1 // this will make the matrix non-singular, if we suppose voltage at node 0 is 0, adding 1 won't change the result
	laplacianMatrix := mat.NewDense(len(adjList), len(adjList), values)
	fmt.Println("Laplacian matrix:", values)
	return laplacianMatrix
}
