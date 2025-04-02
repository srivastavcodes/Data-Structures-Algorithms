package lecture2;

public class Node {
      public Node prev;
      public int value;
      public Node next;

      public Node() {
      }

      public Node(int value) {
            this.prev = null;
            this.value = value;
            this.next = null;
      }

      public Node(Node prev, int value) {
            this.prev = prev;
            this.value = value;
            this.next = null;
      }

      public Node(Node prev, int value, Node next) {
            this.prev = prev;
            this.value = value;
            this.next = next;
      }
}
