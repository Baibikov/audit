# Audit service

### Make log table in clickhouse databse
```sql
create table if not exists log (
    objectCode String,
    actionCode String,
    data Map(String, String),
    createdAt DateTime
)
engine = MergeTree()
    order by createdAt;
```

### Generate grpc server
```bash
make proto
```

### Run application

```bash
go run cmd/app/*.go
```