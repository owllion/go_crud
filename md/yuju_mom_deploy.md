# YUJU MOM deploy 

- 終端執行 deploy.sh後會要你登入: user@36.64.27.229's password:(就是50984878)
- 成功登入後會進到 user@YUJU-ES:(這是直接跑到他們那邊的local去部屬，不用anydesk)
- 然後 cd Desktop -> cd mip(就是MOM，但SAT_A1_MES也是部到 mip)
- 那邊網路很爛，所以在git bash裡面執行 ./deploy.sh 可能會報錯，長這樣:
```Text
Error response from daemon: Get "https://registry-1.docker.io/v2/": dial tcp: lookup registry-1.docker.io: Temporary failure in name resolution

``` 
還有其他不同錯誤，但總之看到這就是他們網路太慢，就繼續執行 ./deploy.sh直到出現綠色的done

- 詳情請見 https://www.notion.so/yuju_mom-ip-e42d92c3063e46ec8bb78418b7ba9a7a
