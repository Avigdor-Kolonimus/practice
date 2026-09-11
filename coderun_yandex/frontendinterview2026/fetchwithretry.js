// https://coderun.yandex.ru/selections/frontend-interview-2026/problems/fetch-with-retry
// FetchWithRetry - problem 1
/**
 * Функция, которая запрашивает данные с помощью fetcher
 * и повторяет запрос в случае ошибки. Запросы повторяются
 * до тех пор, пока не будет получен успешный ответ
 * или пока не будут исчерпаны все попытки.
 * 
 * @param {function} fetcher - функция, которая возвращает Promise
 * @param {number} count - максимальное количество попыток
 */
export async function fetchWithAutoRetry(fetcher, count) {
    try {
     return await fetcher()
    }
    catch(e) {
     if (count === 0) throw e;
     return fetchWithAutoRetry(fetcher, count - 1)
    }
 };