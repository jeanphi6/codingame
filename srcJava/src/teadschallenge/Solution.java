package teadschallenge;

import java.util.*;
import java.io.*;
import java.math.*;
import java.util.stream.Collectors;

/**
 * https://www.codingame.com/ide/puzzle/teads-sponsored-contest
 *
 * Algo :
 * - stockage des noeuds avec une liste de noeud adjacent.
 * - suppression des noeuds ayant un seul lien.
 */
public class Solution {

    private static class Node {
        private int id;
        private Set<Node> nodes;

        public Node(int id) {
            this.id = id;
            this.nodes = new HashSet<Node>();
        }

        public int getId() {
            return this.id;
        }

        public void addNode(Node n) {
            this.nodes.add(n);
        }

        public void remove(Node n) {
            this.nodes.remove(n);
        }

        public void removeFromNodes() {
            this.nodes.stream().forEach(node -> node.remove(this));
        }

        public boolean isLeaf() {
            return nodes.size() == 1;
        }
    }

    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt(); // the number of adjacency relations
        Map<Integer, Node> nodes = new HashMap<Integer, Node>();
        Node tmp = null;
        Node tmp2 = null;

        for (int i = 0; i < n; i++) {
            int xi = in.nextInt(); // the ID of a person which is adjacent to yi
            int yi = in.nextInt(); // the ID of a person which is adjacent to xi

            // create or get node
            tmp = nodes.computeIfAbsent(xi, key -> new Node(key));

            // check and add child node
            tmp2 = nodes.computeIfAbsent(yi, key -> new Node(key));

            // add node links
            tmp2.addNode(tmp);
            tmp.addNode(tmp2);
        }

        // Write an action using System.out.println()
        // To debug: System.err.println("Debug messages...");

        // The minimal amount of steps required to completely propagate the advertisement
        System.out.println(findMinimumPath(nodes));
    }

    /**
     * Find minimum path of the graph.
     *
     * @return the minimum path.
     */
    private static int findMinimumPath(Map<Integer, Node> nodes) {
        int minPath = 0;

        while (nodes.size() > 1) {
            removeLeaves(nodes);
            minPath++;
        }

        return minPath;
    }

    private static void removeLeaves(Map<Integer, Node> nodes) {
        Set<Node> leaves = nodes.values().stream().filter(Node::isLeaf).collect(Collectors.toSet());
        leaves.stream().forEach(leaf -> {
            nodes.remove(leaf.getId());
            leaf.removeFromNodes();
        });
    }
}
