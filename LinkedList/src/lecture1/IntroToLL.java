package lecture1;

public class IntroToLL {

    public static Node constructLL(int[] arr) {
        Node head = new Node(arr[0]);
        Node tail = head;
        for (int i = 1; i < arr.length; i++) {
            Node temp = new Node(arr[i]);
            tail.next = temp;
            tail = temp;
        }
        return head;
    }

    public static int length(Node head) {
        Node temp = head;
        int count = 0;
        while (temp != null) {
            count++;
            temp = temp.next;
        }
        return count;
    }

    public static int checkIfPresent(Node head, int val) {
        Node temp = head;
        while (temp != null) {
            if (temp.value == val) return 1;
            temp = temp.next;
        }
        return 0;
    }

    public static Node removesHead(Node node) {
        if (node == null) return null;
        return node.next;
    }

    public static void printList(Node node) {
        while (node != null) {
            System.out.print(node.value + " ");
            node = node.next;
        }
    }

    public static Node removeTail(Node node) {
        if (node == null || node.next.next == null) return null;
        Node temp = node;
        while (temp.next.next != null) {
            temp = temp.next;
        }
        temp.next = null;
        return node;
    }

    public static Node removeKthNode(Node node, int k) {
        if (node == null) return null;
        if (k == 1) return node.next;
        int count = 0;
        Node temp = node;
        Node prev = null;
        while (temp != null) {
            count++;
            if (count == k) {
                prev.next = prev.next.next;
                break;
            }
            prev = temp;
            temp = temp.next;
        }
        return node;
    }

    public static Node removeElmnt(Node node, int k) {
        if (node == null) return null;
        if (node.value == k) return node.next;
        Node temp = node;
        Node prev = null;
        while (temp != null) {
            if (temp.value == k) {
                prev.next = prev.next.next;
                break;
            }
            prev = temp;
            temp = temp.next;
        }
        return node;
    }
}
