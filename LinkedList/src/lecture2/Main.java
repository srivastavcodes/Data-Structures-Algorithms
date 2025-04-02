package lecture2;

public class Main {
      public static void main(String[] args) {
            int[] arr = {1, 2, 3, 4, 5, 6, 7};

            IntroToDLL dll = new IntroToDLL();
            Node head = dll.convertTODLL(arr);
            dll.printDLL(head);

            System.out.println();

            Node afterDeletion = dll.removeKthNode(head, 0);
            dll.printDLL(afterDeletion);

            System.out.println();

            dll.deleteNode(afterDeletion.next);
            dll.printDLL(afterDeletion);

            System.out.println();

            Node afterHeadInsertion = dll.insertBeforeHead(afterDeletion, 0);
            dll.printDLL(afterHeadInsertion);

            System.out.println();

            Node afterTailInsertion = dll.insertBeforeTail(afterHeadInsertion, 6);
            dll.printDLL(afterTailInsertion);

            System.out.println();

            Node kthIndexInsertion = dll.insertBeforeKthIndex(afterTailInsertion, 11, 4);
            dll.printDLL(kthIndexInsertion);

            System.out.println();

            dll.insertBeforeKthNode(kthIndexInsertion.next.next.next, 57);
            dll.printDLL(kthIndexInsertion);

            System.out.println();
      }
}
