select distinct
    l.n as ConsecutiveNums
from (
    select
        num as n,
        lead(num) over (order by id asc) as n2,
        lead(num, 2) over (order by id asc) as n3
    from Logs
) as l
where
    l.n = l.n2 and l.n = l.n3
;
