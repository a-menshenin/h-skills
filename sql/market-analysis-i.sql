-- https://leetcode.com/problems/market-analysis-i/
select
    u.user_id as buyer_id,
    max(u.join_date) as join_date,
    count(
        case 
        when date_part('year', o.order_date) = 2019
        then o.order_id
        else null
        end
        ) as orders_in_2019
from
    Users u
left join Orders o on o.buyer_id = u.user_id
group by u.user_id
;
