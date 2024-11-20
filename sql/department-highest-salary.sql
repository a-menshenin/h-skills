-- https://leetcode.com/problems/department-highest-salary/
with maxSalary(max_salary, departmentId) as (
    select
        max(e2.salary) as max_salary,
        departmentId
    from Employee e2
    group by
        departmentId
)

select
    d.name as Department,
    e.name as Employee,
    e.salary as Salary
from
    Employee e
join Department d on d.id = e.departmentId
where
    e.salary = (select max_salary from maxSalary m where m.departmentId = e.departmentId)
