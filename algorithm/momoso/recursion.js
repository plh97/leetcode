const obj = {
  "A": 1,
  "B.A": 2,
  "B.B": 3,
  "CC.D.E": 4,
  "CC.D.F": 5
}

function tree(obj) {
  let newTree = {}
  for (let i in obj) {
    find(i.split("."), obj[i], newTree)
  }
  return newTree
}

// output 递归查找函数
function find(keys, val, currentTree) {
  const temp = keys.shift()
  if (keys.length > 0) {
    if (!currentTree[temp]) {
      currentTree[temp] = {}
    }
    find(keys, val, currentTree[temp])
  } else {
    currentTree[temp] = val
  }
}

console.log(
  tree(obj)
);

// { A: 1, B: { A: 2, B: 3 }, CC: { D: { E: 4, F: 5 } } }