package lecture2;

public class Main {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};

        IntroToDLL dll = new IntroToDLL();
        Node head = dll.convertTODLL(arr);
        dll.printDLL(head);
        System.out.println();
        Node afterDeletion = dll.deleteKthNode(head, 0);
        dll.printDLL(afterDeletion);
    }
}
