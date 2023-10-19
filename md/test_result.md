# 測試紀錄

### 以下全都不行

1. X 最一般的clauses
```Golang
	Mock.ExpectQuery(`SELECT * FROM "enrollment"."student" WHERE "id" = $1`).WithArgs(5).WillReturnRows(rows)

```
  - 沒寫WithArgs 沒用
  - 直接寫 5 / ? / $1 全都不行

2. X 正規表達
```Golang
	Mock.ExpectQuery("^SELECT (.+) FROM \"enrollment\".\"student\" WHERE \"id\" = \\$1$").WithArgs(5).WillReturnRows(rows)
```

3. X 只寫 SELECT(.+)，後這代表不限欄位，where也可以寫(.+)

```Golang
	Mock.ExpectQuery("SELECT (.+)").WithArgs(5).WillReturnRows(rows)

```


