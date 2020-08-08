import java.util.Collection;
import java.util.HashSet;

/*
Detect a cycle in a linked list. Note that the head pointer may be 'null' if the list is empty.

A Node is defined as: 
    class Node {
        int data;
        Node next;
    }
*/
public class Program {
  public static void main(String[] args) {

    Node node1 = new Node(1, null);
    Node node2 = new Node(2, null);
    Node node3 = new Node(3, null);
    Node node5 = new Node(5, null);
    Node node4 = new Node(4, null);

    node1.next = node2;
    node2.next = node3;
    node3.next = node4;
    node4.next = node5;
    node5.next = node3;

    System.out.println(hasCycle(node1));
  }

  static boolean hasCycle(Node head) {
      Collection<Integer> nodesSet = new HashSet<Integer>();
      Node node = head; // initializing node aux
      
      if(node == null) // empty list
          return false;
      
      while(!nodesSet.contains(node.hashCode())) {
          if(node.next == null) { // end of list
              return false;
          }
          nodesSet.add(node.hashCode());
          node = node.next;
      }
      
      return true;
  }
}

class Node {
  public int data;
  public Node next;

  public Node(int data, Node next) {
    this.data = data;
    this.next = next;
  }
}