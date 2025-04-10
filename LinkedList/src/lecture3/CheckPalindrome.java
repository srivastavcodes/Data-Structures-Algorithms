package lecture3;

public class CheckPalindrome {

      public static Node reverseLinkedList(Node head) {
            Node prev = null;
            Node curr = head;

            while (curr != null) {
                  Node next = curr.next;
                  curr.next = prev;
                  prev = curr;
                  curr = next;
            }
            return prev;
      }

      @SuppressWarnings("JavaExistingMethodCanBeUsed")
      public static boolean isPalindrome(Node head) {
            if (head == null || head.next == null) return true;

            Node slow = head, fast = head;
            while (fast != null && fast.next != null) {
                  slow = slow.next;
                  fast = fast.next.next;
            }
            Node newHead = reverseLinkedList(slow.next);
            Node first = head, second = newHead;
            while (second != null) {
                  if (second.value != first.value) {
                        reverseLinkedList(newHead);
                        return false;
                  }
                  first = first.next;
                  second = second.next;
            }
            reverseLinkedList(newHead);
            return true;
      }
}
