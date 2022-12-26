```sql
CREATE TABLE table_name (
    field_name TYPE CONSTRAINTS,
    field_name TYPE(args) CONSTRAINTS
);
```

```SQL
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```