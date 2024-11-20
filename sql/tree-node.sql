-- https://leetcode.com/problems/tree-node/description/

-- Непонятно, почему на leetcode не срабатывает условие p_id = null
select
    t.id,
    case 
        when t.p_id IS NULL then 'Root'
        when t.id in (select t.id from Tree t where t.id in (select distinct p_id from Tree)) then 'Inner'
        else 'Leaf'
    end as type
from
    Tree t
;
