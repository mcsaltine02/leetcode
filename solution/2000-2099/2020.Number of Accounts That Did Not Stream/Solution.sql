# Write your MySQL query statement below
SELECT COUNT(DISTINCT sub.account_id) AS accounts_count
FROM
    Subscriptions AS sub
    LEFT JOIN Streams AS ss ON sub.account_id = ss.account_id
WHERE
    YEAR(start_date) <= 2021
    AND YEAR(end_date) >= 2021
    AND (YEAR(stream_date) != 2021 OR stream_date > end_date);
