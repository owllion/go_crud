# gorm automigrate

- 功能非常基本，只能幫你新增和刪除，不會幫你更新
- 假如原本只有id name email，你後來新增 address跟其他好幾個欄位，他會報錯
- 目前確切原因未知，他說id欄位已存在
- 解法: 去db手動drop table再讓gorm去automigrate一次