// const grid = [
//   [0, 0, 0, 0, 0],
//   [0, 0, 0, 0, 0],
//   [0, 0, 0, 0, 0],
//   [0, 0, 0, 0, 0],
//   [0, 0, 0, 0, 0],
// ]




class Spot {
  constructor(i, j, w, h, ctx) {
    this.i = i; // 他们的横向坐标位置 
    this.j = j; // 他们的纵向坐标位置
    this.width = w; // 单个节点的高度
    this.height = h; // 单个节点的宽度
    this.ctx = ctx; // ctx canvas 绘制实例本省
    this.f = 0;
    this.g = 0;
    this.h = 0;
    this.neighbors = [];
    this.prev;
    this.wall = Math.random() < 0.2;

  }
  show(color) {
    if (this.wall) {
      this.ctx.fillStyle = '#000';
    } else {
      this.ctx.fillStyle = color;
    }
    this.ctx.strokeRect(this.i * this.width, this.j * this.height, this.width, this.height)
    this.ctx.fillRect(this.i * this.width, this.j * this.height, this.width, this.height)
  }
  addNeighbors(grid) {
    const i = this.i
    const j = this.j
    if (grid[i + 1]) {
      this.neighbors.push(grid[i + 1][j])
    }
    if (i > 0) {
      this.neighbors.push(grid[i - 1][j])
    }
    if (grid[i][j + 1]) {
      this.neighbors.push(grid[i][j + 1])
    }
    if (j > 0) {
      this.neighbors.push(grid[i][j - 1])
    }
  }
}

const sleep = ms => new Promise(res => setTimeout(res, ms))

class Astar {
  constructor(cols, rows) {
    this.start;
    this.end;
    this.grid = new Array(cols).fill(0).map(e => new Array(rows).fill(0));
    this.cols = cols;
    this.rows = rows;
    this.width = 400;
    this.height = 400;
    this.openSet = []; // 已经访问过的节点 -> 绿色
    this.closedSet = []; // 已经访问过的节点 -> 红色
    this.w = this.width / cols; // 单个格子宽度
    this.h = this.height / rows; // 单个格子高度
    this.path = [];
    this.current;

    console.log(this, "A*")
    const canvas = document.querySelector('canvas')
    const ctx = canvas.getContext('2d')
    // stroke grid
    this.loop((i, j) => {
      this.grid[i][j] = new Spot(i, j, this.w, this.h, ctx);
    })
    this.loop((i, j) => {
      this.grid[i][j].addNeighbors(this.grid)
    })
    // stroke grid
    this.start = this.grid[0][0];
    this.end = this.grid[this.rows - 1][this.cols - 1];
    this.openSet.push(this.start)

    this.timer = setInterval(e => {
      this.draw()
    }, 0)
  }

  loop(callback) {
    for (let i = 0; i < this.cols; i++) {
      for (let j = 0; j < this.rows; j++) {
        callback(i, j)
      }
    }
  }

  heuristic(a, b) {
    var di = b.i - a.i
    var dj = b.j - a.j
    return di * di + dj * dj
  }

  removeFromArray(arr, left) {
    return arr.filter(e => e != left);
  }

  async draw() {
    // for (this.openSet.length > 0;;) {
    if (this.openSet.length > 0) {
      // keep searching going
      let winner = 0;
      this.openSet.forEach((e, i) => {
        if (e.f < this.openSet[winner].f) {
          winner = i;
        }
      })

      this.current = this.openSet[winner]
      if (this.current == this.end) {
        console.log('done')
        clearInterval(this.timer)
      }
      this.openSet = this.removeFromArray(this.openSet, this.current)
      this.closedSet.push(this.current)

      var neighbors = this.current.neighbors;
      for (let i = 0; i < neighbors.length; i++) { // 对当前节点的邻居做遍历循环.如果当前节点
        var neighbor = neighbors[i];
        if (!this.closedSet.includes(neighbor) && !neighbor.wall) { //closedSet 访问过的节点的集合, 如果这个当前节点的邻居是未访问过的节点,继续->
          var tempG = this.current.g + 1 // 得分+1
          if (this.openSet.includes(neighbor)) { // 如果 既是 将要访问的节点, 也是邻居节点
            if (tempG < neighbor.g) { // ???? 如果邻居的 gScore 大于 当前节点 gScore, 给邻居节点得分赋值当前节点
              neighbor.g = tempG; // ???? 意思就是当前邻居节点它之前的得分是更大的, 直接更换成从当前节点走过去会更好
            }
          } else {
            neighbor.g = tempG; // 当前邻居节点是新节点,它的得分直接就是 当前节点+1 因为目前位置,暂时未发现任何去到这个邻居节点的方式,除了当前节点能够访问到.
            this.openSet.push(neighbor) // 将这个邻居节点push进去
          }
          neighbor.h = this.heuristic(neighbor, this.end) // 当前邻居节点的h值就是当前邻居节点到终点的直线距离
          neighbor.f = neighbor.g + neighbor.h // f 就是 物理距离 + g 也就是我所话费的步骤, 也就是已知假定总距离/
          neighbor.prev = this.current;
        }
      }
    } else {
      console.log('no solution')
      // no solution
      clearInterval(this.timer)
    }
    // }
    // background(0);
    this.loop((i, j) => {
      this.grid[i][j].show('rgb(255,255,255)');
    })
    this.closedSet.forEach(e => {
      e.show('rgb(255,0,0)')
    })
    this.openSet.forEach(e => {
      e.show('rgb(0,255,0)')
    })
    // if (this.current == this.end) {
    let temp = this.current;
    this.path = [temp];
    while (temp.prev) {
      this.path.push(temp.prev)
      temp = temp.prev
    }
    this.path.forEach(e => {
      e.show('blue')
    })
    // }
  }
}

new Astar(50, 50)
// astar.draw()