// time O(n) 只能有一层for循环
// space O(1) -> 意思就是只能采用 交换指针策略,才能保证不会出现多余内存占用

// 采用两个指针策略,
// 不需要保持原有顺序

const sort = arr => {
  var i = 0;
  var j = arr.length - 1;
  while (i < j) {
    while (arr[i] > 0) {
      i++
    }
    while (arr[j] < 0) {
      j--
    }
    [arr[i], arr[j]] = [arr[j], arr[i]]
    i++;
    j--
  }
  return arr;
}

const ints = [6, 4, -3, 5, -2, -1, 0, 1, -9]

console.log(
  sort(ints)
);
// [ 6, 4, 1, 5, 0, -1, -2, -3, -9 ] 