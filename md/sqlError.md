# ERROR: invalid input syntax for type double precision: ""

- 有這錯是因為用SPLIT_PART轉換出來的是空字串，此時又用CAST就會報這錯，會產生 result.Error, 會造成500


- 原本長這樣，會error: CAST(SPLIT_PART(r10."PN", '-', 4) AS double precision)，改成下面這樣才ok，會判斷當前split的部分是否不等於(就是<>)""，如果不是的話才cast，反之就直接變成null between ? AND ? 
```Golang
	condition := `
		CASE 
			WHEN TRIM(SPLIT_PART(r10."PN", '-', 4)) <> '' THEN 
				CAST(SPLIT_PART(r10."PN", '-', 4) AS double precision) 
			ELSE 
				NULL 
		END BETWEEN ? AND ?
    `
```

- 而且要先判斷result.Error(原本是result.RowsAffected先寫，但要是中途有錯誤，就會是404，但其實我目前這情況是"失敗"，不是"查不到"! 差很多!)

```Golang
	if result.Error != nil {
			// util.Log("查詢庫存資料失敗", nil, result.Error.Error())
			g.SendResponse(500, "查詢庫存資料失敗", result.Error.Error())
			return
		}

	if result.RowsAffected == 0 {
		g.SendResponse(404, "查無庫存資料", res)
		return
	}

```
   