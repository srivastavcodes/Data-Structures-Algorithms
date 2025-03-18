package lecture1;

public class IntroToLL {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        Node node = constructLL(arr);
        node = removeKthNode(node, 8);
        printList(node);
        System.out.println();
        node = removeElmnt(node, 2);
        printList(node);
    }

    private static Node constructLL(int[] arr) {
        Node head = new Node(arr[0]);
        Node tail = head;
        for (int i = 1; i < arr.length; i++) {
            Node temp = new Node(arr[i]);
            tail.next = temp;
            tail = temp;
        }
        return head;
    }

    private static int length(Node head) {
        Node temp = head;
        int count = 0;
        while (temp != null) {
            count++;
            temp = temp.next;
        }
        return count;
    }

    private static int checkIfPresent(Node head, int val) {
        Node temp = head;
        while (temp != null) {
            if (temp.data == val) return 1;
            temp = temp.next;
        }
        return 0;
    }

    private static Node removesHead(Node node) {
        if (node == null) return null;
        return node.next;
    }

    private static void printList(Node node) {
        while (node != null) {
            System.out.print(node.data + " ");
            node = node.next;
        }
    }

    private static Node removeTail(Node node) {
        if (node == null || node.next.next == null) return null;
        Node temp = node;
        while (temp.next.next != null) {
            temp = temp.next;
        }
        temp.next = null;
        return node;
    }

    private static Node removeKthNode(Node node, int k) {
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

    private static Node removeElmnt(Node node, int k) {
        if (node == null) return null;
        if (node.data == k) return node.next;
        Node temp = node;
        Node prev = null;
        while (temp != null) {
            if (temp.data == k) {
                prev.next = prev.next.next;
                break;
            }
            prev = temp;
            temp = temp.next;
        }
        return node;
    }
}

class Node {
    public int data;
    public Node next;

    public Node() {
        this.data = 0;
        this.next = null;
    }

    public Node(int data) {
        this.data = data;
        this.next = null;
    }

    public Node(int data, Node next) {
        this.data = data;
        this.next = next;
    }
}
