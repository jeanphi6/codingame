package teadschallenge

import "fmt"
import "os"

/**
 * https://www.codingame.com/ide/puzzle/teads-sponsored-contest
 *
 * Algo :
 * - stockage des noeuds avec une liste de noeud adjacent.
 * - suppression des noeuds ayant un seul lien.
 */

func main() {
	// n: the number of adjacency relations
	var n int
	fmt.Scan(&n)

	var nodes []node = make([]node, n)

	for i := 0; i < n; i++ {
		// xi: the ID of a person which is adjacent to yi
		// yi: the ID of a person which is adjacent to xi
		var xi, yi int
		fmt.Scan(&xi, &yi)

		tmp := getOrCreateNode(nodes, xi)
		tmp2 := getOrCreateNode(nodes, yi)

		nodes = append(nodes, tmp)
		nodes = append(nodes, tmp2)

		tmp.nodes = append(tmp.nodes, tmp2)
		tmp2.nodes = append(tmp2.nodes, tmp)

	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// The minimal amount of steps required to completely propagate the advertisement
	result := findMinimumPath(nodes)
	fmt.Println(result)

}

func findMinimumPath(nodes []node) int {
	var result int = 0

	for len(nodes) > 1 {
		fmt.Fprint(os.Stderr, "-> ")
		fmt.Fprintln(os.Stderr, len(nodes))
		removeLeaves(nodes)
		result ++
	}

	return result
}
func removeLeaves(nodes []node) {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].isLeaf() {
			nodes[i] = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			nodes[i].removeFromNodes()
		}
	}
}

type node struct {
	id    int
	nodes []node
}

func (n node) addNode(nToAdd node) node {
	n.nodes = append(n.nodes, nToAdd)
	return n
}

func (n node) removeNode(nToRemove node) node {
	for i := 0; i < len(n.nodes); i ++ {
		if n.nodes[i].id == nToRemove.id {
			n.nodes[i] = n.nodes[len(n.nodes)-1]
			//n.nodes[len(n.nodes)-1] = nil
			n.nodes = n.nodes[:len(n.nodes)-1]
		}
	}
	return n
}

func (n node) removeFromNodes() {
	for i := 0; i < len(n.nodes); i ++ {
		for j := 0; j < len(n.nodes[i].nodes); j++ {
			if n.nodes[i].nodes[j].id == n.id {
				n.nodes[i].nodes[j] = n.nodes[i].nodes[len(n.nodes)-1]
				//n.nodes[len(n.nodes)-1] = nil
				n.nodes[i].nodes = n.nodes[i].nodes[:len(n.nodes[i].nodes)-1]
			}
		}
	}
}

func (n node) isLeaf() bool {
	return len(n.nodes) == 1
}

func getOrCreateNode(nodes []node, valueToFind int) node {
	var result node
	var isFind bool = false
	var tabs []node = make([]node, 0)

	for i := 0; i < len(nodes); i ++ {
		if nodes[i].id == valueToFind {
			result = nodes[i]
			isFind = true
		}
	}

	if !isFind {
		result = node{valueToFind, tabs}
	}

	return result
}

