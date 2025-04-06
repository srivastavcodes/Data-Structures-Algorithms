package lecture3;

import static lecture3.Node.convertTODLL;
import static lecture3.Node.printDLL;

// Tortoise and hare algorithm
public class MiddleOfLinkedList {

      public static void main(String[] args) {
            int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9};
            Node node = convertTODLL(arr);
            printDLL(node);
            System.out.println();
            System.out.println(findMiddle(node).value);
      }

      public static Node findMiddle(Node head) {
            Node slow = head, fast = head;

            while (fast != null && fast.next != null) {
                  slow = slow.next;
                  fast = fast.next.next;
            }
            return slow;
      }
}
