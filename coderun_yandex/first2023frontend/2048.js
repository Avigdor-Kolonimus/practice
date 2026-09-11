// https://coderun.yandex.ru/selections/first-2023-frontend/problems/2048
// 2048 - problem 1
module.exports = function solution(field, moves) {
    moves = moves.split(" ");
  
    moves.forEach((m) => {
      if (m === "U") {
        moveUp(field);
      }
  
      if (m === "D") {
        moveDown(field);
      }
  
      if (m === "L") {
        moveLeft(field);
      }
  
      if (m === "R") {
        moveRight(field);
      }
    });
  
    return field;
  };
  
  function move(arr) {
    shift(arr);
    sum(arr);
    shift(arr);
  }
  
  function shift(arr) {
    let offset = 0;
    for (let i = 0; i < 4; i++) {
      if (arr[i] === 0) {
        offset++;
      }
      if (arr[i - offset] === 0) {
        arr[i - offset] = arr[i];
        arr[i] = 0;
      }
    }
  }
  
  function sum(arr) {
    for (let i = 0; i < 4; i++) {
      if (arr[i] !== 0 && arr[i] === arr[i + 1]) {
        arr[i] *= 2;
        arr[i + 1] = 0;
        i++;
      }
    }
  }
  
  function moveUp(field) {
    const w = field[0].length;
    const h = field.length;
  
    for (let x = 0; x < w; x++) {
      const line = [];
  
      for (let y = 0; y < h; y++) {
        line.push(field[y][x]);
      }
  
      move(line);
  
      for (let y = 0; y < h; y++) {
        field[y][x] = line[y];
      }
    }
  }
  
  function moveDown(field) {
    const w = field[0].length;
    const h = field.length;
  
    for (let x = 0; x < w; x++) {
      const line = [];
  
      for (let y = h - 1; y >= 0; y--) {
        line.push(field[y][x]);
      }
  
      move(line);
  
      for (let y = h - 1; y >= 0; y--) {
        field[y][x] = line[h - 1 - y];
      }
    }
  }
  
  function moveLeft(field) {
    const w = field[0].length;
    const h = field.length;
  
    for (let y = 0; y < h; y++) {
      const line = [];
  
      for (let x = 0; x < w; x++) {
        line.push(field[y][x]);
      }
  
      move(line);
  
      for (let x = 0; x < w; x++) {
        field[y][x] = line[x];
      }
    }
  }
  
  function moveRight(field) {
    const w = field[0].length;
    const h = field.length;
  
    for (let y = 0; y < h; y++) {
      const line = [];
  
      for (let x = w - 1; x >= 0; x--) {
        line.push(field[y][x]);
      }
  
      move(line);
  
      for (let x = w - 1; x >= 0; x--) {
        field[y][x] = line[w - 1 - x];
      }
    }
  }