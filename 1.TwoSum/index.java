class Solution {
  public int[] twoSum(int[] nums, int target) {
    for (int i = 0; i < nums.length; i++) {
      for (int j = i+1; j < nums.length; j++) {
        if(nums[j]+ nums[i] == target){
          return new int[] {i,j};
        }
      }
    }
    throw new Error("not found");
  };
}



public class index {
  /* 第一个Java程序
   * 它将打印字符串 Hello World
   */
  public static void main(String []args) {
    int[] nums = new int[3];         // 首选的方法
    nums[0] = 3;
    nums[1] = 2;
    nums[2] = 4;
    Solution solution = new Solution();




    int[] res = solution.twoSum(nums,6);
    System.out.print(res);
  }
}
