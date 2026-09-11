// https://coderun.yandex.ru/selections/frontend-interview-2026/problems/memoize
// Memoize - problem 5
export function memoize(fn) {
    const cache = new Map

   return function (...args) {
       let key = JSON.stringify(args)

       if (cache.has(key)) {
           return cache.get(key)
       }

       const result = fn(...args)

       cache.set(key, result)
       
       return  result
   }
}