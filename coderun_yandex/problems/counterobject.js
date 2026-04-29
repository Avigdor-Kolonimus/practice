/**
 * Реализуйте функцию, которая создаёт объект счётчика
 * с базовыми методами: increment, decrement, reset.
 * 
 * @param {number} init - начальное значение счётчика
 * @returns {object} объект с методами increment, decrement, reset
 */
export function createCounter(init) {
    let value = init;

    return {
        increment() {
            value += 1;
            return value;
        },
        decrement() {
            value -= 1;
            return value;
        },
        reset() {
            value = init;
            return value;
        }
    }
}; 