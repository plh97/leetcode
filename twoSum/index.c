#include <stdio.h>
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *twoSum(int *nums, int numsSize, int target)
{
  static int res[2]={};
  for (int i = 0; i < numsSize - 1; i++)
  {
    /* code */
    for (int j = i+1; j < numsSize; j++)
    {
      if(nums[i]+nums[j] == target){
        res[0] = i;
        res[1] = j;
      }
    };
  };
  return res;
}

int main(int argc, char const *argv[])
{
  int nums[3];
  nums[0] = 3;
  nums[1] = 2;
  nums[2] = 4;

  int *res = twoSum(nums, 3, 6);
  printf("%n", res);
  return 0;
}
