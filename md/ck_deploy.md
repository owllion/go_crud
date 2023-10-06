# CK部屬流程

-------------------- local --------------------------------------------
- git pull
- ./deploy.sh
- Git Bash 再次 ./deploy.sh
- wait 
- 確認 docker pushed 成功 -> 要確認waiting的部分是否都已經pushed或是否是already exists,如果都沒有就是push失敗，要一直重新執行deploy

-------------------- CK部屬 -----------------------------------------
- 開 AnyDesk 連線到 CK 主機 (無外網/密碼:satsys+統編)
- 開終端執行 ./deploy.sh
- 確認有拿到更新過的docker image -> 如果顯示 Image is up to date for xxx/xxx 就是失敗(沒有拿到更新過的)，要重新去local部屬
- 部屬成功
