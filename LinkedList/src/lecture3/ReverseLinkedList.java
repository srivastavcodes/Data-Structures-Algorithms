package lecture3;

import static lecture3.Node.convertTODLL;
import static lecture3.Node.printDLL;

public class ReverseLinkedList {
      public static void main(String[] args) {
            int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9};
            Node node = convertTODLL(arr);
            printDLL(node);
            System.out.println();
            printDLL(reverseLinkedList(node));
      }

      public static Node reverseLinkedList(Node head) {
            if (head == null || head.next == null) {
                  return head;
            }
            Node newHead = reverseLinkedList(head.next);
            Node front = head.next;
            front.next = head;
            head.next = null;
            return newHead;
      }
}
