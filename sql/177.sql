
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
      SET N := N-1;
RETURN (
    # Write your MySQL query statement below.
    select ifNull (
    (select distinct salary
    from Employee
    order by salary DESC
    limit N, 1), null)
    as SecondHighestSalary
    );
END
/*使用distinct 成绩进行成绩去重*/
/*ifNull函数判断*/