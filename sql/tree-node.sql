-- https://leetcode.com/problems/tree-node/description/

select
    t.id,
    case 
        when t.p_id = null then 'Root'
        when t.id in (select t.id from Tree t where t.id in (select distinct p_id from Tree)) then 'Inner'
        else 'Leaf'
    end as type
from
    Tree t
;
