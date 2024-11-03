update Salary as s set
  sex = s2.newsex
from (values
  ('m', 'f'),
  ('f', 'm')
) as s2(sex, newsex)
where s.sex = s2.sex
;
