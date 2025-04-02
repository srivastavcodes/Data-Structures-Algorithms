package lecture2;

public class IntroToDLL {

      public Node convertTODLL(int[] arr) {
            Node head = new Node(arr[0]);
            Node prev = head;
            for (int i = 1; i < arr.length; i++) {
                  Node curr = new Node(prev, arr[i]);
                  prev.next = curr;
                  prev = curr;
            }
            return head;
      }

      public void printDLL(Node dll) {
            while (dll != null) {
                  System.out.print(dll.value + " ");
                  dll = dll.next;
            }
      }

      public Node removeKthNode(Node head, int k) {
            if (head == null) return null;
            Node curr = head;
            int count = 0;

            while (curr.next != null) {
                  count++;
                  if (count == k) break;
                  curr = curr.next;
            }
            Node back = curr.prev;
            Node front = curr.next;

            if (back == null && front == null) {
                  return null;
            } else if (back == null) {
                  return removeHead(head);
            } else if (front == null) {
                  return removeTail(head);
            }
            back.next = front;
            front.prev = back;

            curr.next = null;
            curr.prev = null;

            return head;
      }

      public Node removeHead(Node head) {
            if (head == null || head.next == null) return null;
            Node curr = head;
            head = head.next;
            head.prev = null;
            curr.next = null;
            return head;
      }

      public Node removeTail(Node head) {
            if (head == null || head.next == null) return null;
            Node curr = head;
            while (curr.next != null) {
                  curr = curr.next;
            }
            Node tail = curr.prev;
            tail.next = null;
            curr.prev = null;
            return head;
      }

      public void deleteNode(Node node) {
            Node back = node.prev;
            Node front = node.next;

            if (front == null) {
                  back.next = null;
                  node.prev = null;
                  return;
            }
            back.next = front;
            front.prev = back;

            node.next = node.prev = null;
      }
}
