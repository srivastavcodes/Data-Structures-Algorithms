package lecture1;

public class IntroToLL {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        Node head = constructLL(arr);
        Node temp = head;
        while (temp != null) {
            System.out.print(temp.data+ " ");
            temp = temp.next;
        }
        System.out.print("\n" + length(head));
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
