import java.io.*;
import java.lang.Math;
import java.util.ArrayList;
import java.util.Map;
import java.util.stream.Collectors;

class Day1 {
    public Day1(String fileName) throws IOException {
        String line;
        BufferedReader br = new BufferedReader(new FileReader(fileName));
        while ((line = br.readLine()) != null) {
            String[] nums = line.split("   ");
            fList1.add(Integer.decode(nums[0]));
            fList2.add(Integer.decode(nums[1]));
        }
        br.close();
    };

    private ArrayList<Integer> fList1 = new ArrayList<>();

    private ArrayList<Integer> fList2 = new ArrayList<>();

    public Integer solvePuzzle1() {
        Integer sum = 0;
        fList1.sort(null);
        fList2.sort(null);
        for (int i = 0; i < fList1.size(); i++){
            sum += Math.abs(fList2.get(i) - fList1.get(i));
        }
        return sum;
    }

    public Integer solvePuzzle2() {
        Integer similarity = 0;

        Map<Integer, Long> map = fList2.stream().collect(Collectors.groupingBy(i -> i, Collectors.counting()));

        for (int i = 0; i < fList1.size(); i++) {
            Long val = map.get(fList1.get(i));
            if (val != null) {
                similarity += fList1.get(i) * val.intValue();
            }
            
        }
        
        return similarity;
    }

    public static void main(String[] args) throws IOException
    {
        Day1 solver = new Day1("Day1\\input.csv");
        
        Integer puzzleOne = solver.solvePuzzle1();
        System.out.println("The distance between the two lists is: " + puzzleOne);

        Integer puzzleTwo = solver.solvePuzzle2();
        System.out.println("The similarity of the lists is: " + puzzleTwo);
    }
}
