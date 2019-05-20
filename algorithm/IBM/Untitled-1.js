// 计算小于n的所有质数
function prime(n) {
  res = [];
  for (let i = 3; i < n; i++) {
    if (isPrime(i)) {
      res.push(i)
    }
  }
  return res
}

function isPrime(n){
  for (let i = 2; i < n/2+1; i++) {
    if (n%i === 0) {
      return false
    } 
  }
  return true
}

console.log(
    prime(121),
    prime(121).length
)