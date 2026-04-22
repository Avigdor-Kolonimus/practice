/**
* @param {number} N - целое число, количество сотрудников готовых к объединению
* @param {number[]} staff - массив длины N с грейдами доступных сотрудников
* @param {number} K - целое число, количество доступных клавиатур
* @returns {number}
*/
module.exports = function (N, staff, K) {
    const cnt = Array(26).fill(0);

    for (let x of staff) {
        cnt[x]++;
    }

    let sum = 0;

    for (let grade = 25; grade >= 0 && K > 0; grade--) {
        let take = Math.min(cnt[grade], K);
        sum += take * grade;
        K -= take;
    }

    return sum;
}