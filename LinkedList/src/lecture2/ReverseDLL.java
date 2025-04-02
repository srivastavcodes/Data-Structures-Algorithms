package lecture2;

public class ReverseDLL {
      public static void main(String[] args) {
            int[] arr = {1, 2, 3, 4, 5, 6, 7};

            IntroToDLL dll = new IntroToDLL();
            Node head = dll.convertTODLL(arr);
            dll.printDLL(head);

            System.out.println();

            ReverseDLL reverse = new ReverseDLL();
            Node afterReversal = reverse.reverseDLL(head);
            dll.printDLL(afterReversal);

            System.out.println();
      }

      public Node reverseDLL(Node head) {
            if (head == null || head.next == null) {
                  return head;
            }
            Node back = null;
            Node curr = head;
            while (curr != null) {
                  back = curr.prev;
                  curr.prev = curr.next;
                  curr.next = back;
                  curr = curr.prev;
            }
            return back.prev;
      }
}
