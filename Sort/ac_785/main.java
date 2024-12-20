import java.util.Scanner;

public class Main {
    public static void main(String[] agrs) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        
        int[] arr = new int[n];
        
        for(int i = 0; i < n; i++) {
            arr[i] = in.nextInt();
        }
        
        quickSort(arr, 0, n-1);
        
        
        for(int i = 0; i < n; i++) {
            System.out.printf("%d ", arr[i]);
        }
        
        
    }
    
    public static void quickSort(int[] arr, int start, int end) {
        if(start >= end) {
            return;
        }
        int j = partition(arr, start, end);
        quickSort(arr, start, j);
        quickSort(arr, j+1, end);
    }
    
    
    public static int partition(int[] arr, int start, int end) {
        int x = arr[start+(end-start)/2];
        int i = start-1;
        int j = end+1;
        
        while(true) {
            while(i < end && arr[++i] < x);
            while(j > start && arr[--j] > x);
            if(i >= j) {
                break;
            }
            swap(arr, i, j);
        }
        return j;
    }
    
    public static void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }
    
}