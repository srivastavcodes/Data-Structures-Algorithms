package lecture1;

import static lecture1.IntroToLL.*;

public class Main {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9};
        Node node = constructLL(arr);
        node = removeKthNode(node, 8);
        printList(node);
        System.out.println();
        node = removeElmnt(node, 2);
        printList(node);
    }
}
