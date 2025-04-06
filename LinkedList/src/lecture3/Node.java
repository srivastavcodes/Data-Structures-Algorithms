package lecture3;

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

      public static Node convertTODLL(int[] arr) {
            Node head = new Node(arr[0]);
            Node prev = head;
            for (int i = 1; i < arr.length; i++) {
                  Node curr = new Node(prev, arr[i]);
                  prev.next = curr;
                  prev = curr;
            }
            return head;
      }

      public static void printDLL(Node dll) {
            while (dll != null) {
                  System.out.print(dll.value + " ");
                  dll = dll.next;
            }
      }
}
