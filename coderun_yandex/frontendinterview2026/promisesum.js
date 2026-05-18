/**
 * Реализуйте функцию, которая принимает на вход два объекта Promise
 * с типом `number` и возвращает Promise с их суммой
 */
export const addTwoPromises = async function (promise1, promise2) {
    const [v1, v2] = await Promise.all([
    promise1.catch(e => e),
    promise2.catch(e => e),
  ]);

  return v1 + v2;
};