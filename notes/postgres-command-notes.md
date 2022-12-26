
Dropping a Table
```sql
    DROP TABLE IF EXiSTS users;
```

Creating a table
```sql 
    create table users (
     id SERIAL PRIMARY KEY,
     age INT,
     first_name TEXT,
     last_name TEXT,
     email TEXT UNIQUE NOT NULL
);
```

Inserting Values

```sql
INSERT INTO users VALUES(1,29,'Paul', 'Santiago', 'paul@gmail.com');
```
    INSERT INTO users VALUES(1,29,'Paul', 'Santiago', 'paul@gmail.com');
    INSERT 0 1     
    lenslockedv2=# select * from users;
    id | age | first_name | last_name |     email      
    ----+-----+------------+-----------+----------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com 
OR

```sql
INSERT INTO users(age,email,first_name,last_name)
VALUES(30,'francis@gmail.com','Francis','Santiago')
```

    insert into users(age,email,first_name,last_name)
    values(30,'francis@gmail.com','Francis','Santiago');
    INSERT 0 1
    lenslockedv2=# select * from users;
    id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | Francis    | Santiago  | francis@gmail



Note: if you use serial and an insert fail it'll skip the next auto increment value and go on to the next.
 Ex: So if inserting the value for francis fails the first time and its id was 3, but succeeds the second time it skips from entering 3 as the id to 4

        insert into users(age,email,first_name,last_name)
        values(30,'francis@gmail.com','Francis','Santiago');
        ERROR:  duplicate key value violates unique constraint "users_email_key"
        DETAIL:  Key (email)=(francis@gmail.com) already exists.
        lenslockedv2=# insert into users(age,email,first_name,last_name)
        values(30,'francis2@gmail.com','Francis','Santiago');
        INSERT 0 1
        lenslockedv2=# select * from users;
        id | age | first_name | last_name |       email        
        ----+-----+------------+-----------+--------------------
        1 |  29 | Paul       | Santiago  | paul@gmail.com
        2 |  30 | Francis    | Santiago  | francis@gmail.com
        4 |  30 | Francis    | Santiago  | francis2@gmail.com

 -- You can do comments like this

 Query
 
 ```sql
SELECT * FROM users; -- select everything from the users table
SELECT id, email FROM users; --select some fields from the users table
```
select * from users;
 id | age | first_name | last_name |       email        
----+-----+------------+-----------+--------------------
  1 |  29 | Paul       | Santiago  | paul@gmail.com
  2 |  30 | Francis    | Santiago  | francis@gmail.com
  4 |  30 | Francis    | Santiago  | francis2@gmail.com

select age,first_name from users;
 age | first_name 
-----+------------
  29 | Paul
  30 | Francis
  30 | Francis
(3 rows)

Filtering Queries

```sql
SELECT * FROM users WHERE email='paul@gmail.io';
SELECT * FROM users WHERE age < 22;
SELECT * FROM users
WHERE age < 30 OR last_name = 'Santiago' LIMIT 1;
```

    select * from users where email='paul@gmail.com';
    id | age | first_name | last_name |     email      
    ----+-----+------------+-----------+----------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com

    No result example:
    select * from users where age < 22;
    id | age | first_name | last_name | email 
    ----+-----+------------+-----------+-------
    (0 rows) 


    select * from users where email='paul@gmail.com';
    id | age | first_name | last_name |     email      
    ----+-----+------------+-----------+----------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com


Updating Records

```sql
    UPDATE users SET first_name ='John', email='jon@gmail.com' where id=4;
```

EX

    BEFORE: 
    id | age | first_name | last_name |       email        
    ----+-----+------------+-----------+--------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | Francis    | Santiago  | francis@gmail.com
    4 |  30 | Francis    | Santiago  | francis2@gmail.com

    AFTER:
    id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | Francis    | Santiago  | francis@gmail.com
    4 |  30 | John       | Santiago  | jon@gmail.com

Multiple Records at Once

```sql
    UPDATE users SET first_name='IsThirty' where age = 30;
 ```

    BEFORE:
    id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | Francis    | Santiago  | francis@gmail.com
    4 |  30 | John       | Santiago  | jon@gmail.com

    AFTER:
     id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | IsThirty   | Santiago  | francis@gmail.com
    4 |  30 | IsThirty   | Santiago  | jon@gmail.com

Deleting a Record

```sql
    DELETE FROM users where id=1;
```

BEFORE:
    id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    1 |  29 | Paul       | Santiago  | paul@gmail.com
    2 |  30 | Francis    | Santiago  | francis@gmail.com
    4 |  30 | John       | Santiago  | jon@gmail.com

AFTER:

 id | age | first_name | last_name |       email       
----+-----+------------+-----------+-------------------
  2 |  30 | IsThirty   | Santiago  | francis@gmail.com
  4 |  30 | IsThirty   | Santiago  | jon@gmail.com

**Note**
 You won't get the old ID back if you insert a new record

 Ex:

    insert into users(age,email,first_name,last_name)
    values(30,'paul@gmail.com','Paul','Santiago');
    INSERT 0 1
    lenslockedv2=# select * from users;
    id | age | first_name | last_name |       email       
    ----+-----+------------+-----------+-------------------
    2 |  30 | IsThirty   | Santiago  | francis@gmail.com
    4 |  30 | IsThirty   | Santiago  | jon@gmail.com
    5 |  30 | Paul       | Santiago  | paul@gmail.com
