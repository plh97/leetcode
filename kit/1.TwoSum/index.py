class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(0, len(nums)):
          for j in range(i, len(nums)):
            if i>=j:
              continue
            if nums[i] + nums[j] == target:
              return [i,j]
        return 0

a = Solution()

print(a.twoSum(
  [3,2,4],6
))