/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i+1; j < nums.length; j++) {
      if(nums[i] + nums[j] == target){
        return [i,j]
      }
    }
  }
  return '无解';
};

let res = twoSum([3, 2, 3], 6) 
console.log(res)