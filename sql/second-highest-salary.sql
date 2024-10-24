-- Не проходит тест, когда всего две записи и у обоих одинаковое значение
SELECT
    CASE
        WHEN (SELECT COUNT(*) FROM Employee) < 2 
        THEN NULL
        ELSE (
            SELECT salary FROM Employee
            ORDER BY salary ASC
            LIMIT 1
            OFFSET 1
        )
    END AS SecondHighestSalary
;
