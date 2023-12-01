### 要新增的dealPN query
```sql

-- 共模


-- T507061/7053
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T507061%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T507053%';

-- T507075/8202
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T507075%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508202%';

select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508588%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508585%';

-- 58731/32-0D060-C0
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%58732-0D060-C0%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%58731-0D060-C0%';

-- T508172/70
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508170%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508172%';

-- 55063-S8806/7/8
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%55063-S8806%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%55063-S8807%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%55063-S8808%';


select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T656935%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T662675%';

select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508498%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%T508491%';

--有共模 但目前60沒資料 
-- TSPT12060/63
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%TSPT12063%';

--這個適從36-47
-- tspt12036-47
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%tspt1203%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%tspt1204%';

--從64-69
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%TSPT1206%';


-- 1. JK245460-1690/0051(要去掉中間-)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%JK245460%';

-- 2. tspt12078-2087(全都有)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%tspt1207%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%tspt1208%';

-- 3.55660-S8038/32(38沒資料)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%55660-S8032%';

--4. 55660-S8039/37(39沒資料)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%55660-S8037%';

--55660-S8035 / 36(35沒資料)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%55660-S8036%';

--5. 77130/135-3M3-T010-M1
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%771303M3T010M1%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%771353M3T010M1%';

--6. ER3030996/7ZNP2
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%ER3030996ZNP2%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%ER3030997ZNP2%';

--7.ACXE26-02340/02350
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%ACXE26-02350%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%ACXE26-02340%';

--8.TSPT12022-27(都有)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%tspt1202%';


--9. T656939/T666346
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T666346%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T656939%';

--10.T508501/T662687
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T662687%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T508501%';

-- 11.T508177/T662681
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T508177%';

select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%T662681%';



--12.79952-bz120/bz090-c0(090無資料)
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%79952-BZ120-C0%';


--13.55885-KK030/KK040
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%55885-KK040%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" 
ILIKE '%55885-KK030%';



-- 非共模但有查到
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%jk245450-1160%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%116250-3512%';
select * from "0ADM_2BDM".bdm08deal where "dealPN" ILIKE '%jk116761-0641%';

```

```sql

INSERT INTO "1WOS_1RUN".runz0temp 
("PN", "dealPN", "pdtNM", "custID", "custNM", "is_split", "created_at")
SELECT 
  "1WOS_1RUN".temp_bdm08deal."PN",
  "1WOS_1RUN".temp_bdm08deal."dealPN",
  "1WOS_1RUN".temp_bdm08deal."dealNM",
  "1WOS_1RUN".temp_bdm08deal."dealID",
  "1WOS_1RUN".temp_bdm08deal."dealalias",
   true, -- 固定為 true
  NOW() -- 現
FROM 
  "1WOS_1RUN".temp_bdm08deal
WHERE 
  "1WOS_1RUN".temp_bdm08deal."dealPN" ILIKE ANY(ARRAY[
    '%T507061%',
    '%T507053%',
    '%T507075%',
    '%T508202%',
    '%T508588%',
    '%T508585%',
    '%58732-0D060-C0%',
    '%58731-0D060-C0%',
    '%T508170%',
    '%T508172%',
    '%jk116761-0641%',
    '%55063-S8806%',
    '%55063-S8807%',
    '%55063-S8808%',
    '%T656935%',
    '%T662675%',
    '%T508498%',
    '%T508491%',
    '%TSPT12063%',
	  
	 '%tspt12036%',
    '%tspt12037%',
    '%tspt12038%',
    '%tspt12039%',
    '%tspt12040%',
    '%tspt12041%',
    '%tspt12042%',
    '%tspt12043%',
    '%tspt12044%',
    '%tspt12045%',
    '%tspt12046%',
    '%tspt12047%',
	  
    '%TSPT12064%',
    '%TSPT12065%',
    '%TSPT12066%',
    '%TSPT12067%',
    '%TSPT12068%',
    '%TSPT12069%',
	  
    '%TSPT12078%',
    '%TSPT12079%',
    '%TSPT12080%',
    '%TSPT12081%',
    '%TSPT12082%',
    '%TSPT12083%',
    '%TSPT12084%',
    '%TSPT12085%',
    '%TSPT12086%',
    '%TSPT12087%',
	  
    '%JK2454601690%',
	'%JK2454600051%',
	  
	'%TSPT12022%',
    '%TSPT12023%',
    '%TSPT12024%',
    '%TSPT12025%',
    '%TSPT12026%',
    '%TSPT12027%',
	  
    '%jk245450-1160%',
    '%116250-3512%',
    '%jk116761-0641%'
	  
	  
    '%55660-S8032%',
    '%55660-S8037%',
    '%55660-S8036%',
    '%771303M3T010M1%',
    '%771353M3T010M1%',
    '%ER3030996ZNP2%',
    '%ER3030997ZNP2%',
    '%ACXE26-02350%',
    '%ACXE26-02340%',
	  
    '%tspt1202%',
	  
    '%T666346%',
    '%T656939%',
    '%T662687%',
    '%T508501%',
    '%T508177%',
    '%T662681%',
    '%79952-BZ120-C0%',
    '%55885-KK040%',
    '%55885-KK030%'
  ]);


```

0B-PLY-030-1220-2440-HB-TH-F -> 0B-PLY-030-1220-2440-BT-TH-F
0B-PLY-120-1220-2440-HB-TH-F -> 0B-PLY-120-1220-2440-BT-TH-F
0B-PLY-150-1220-2440-HB-TH-F -> 0B-PLY-150-1220-2440-BT-TH-F
0B-PLY-050-1220-2440-DH-TH-F -> 0B-PLY-050-1220-2440-HB-TH-F
(這有兩筆，一筆先改成HB，第2比)
0B-PLY-050-1220-2440-HB-TH-F -> 0B-PLY-050-1220-2440-BT-TH-F


<!-- "0ADM_2BDM".bdm00pdt  -->0B-PLY-050-1220-2440-HB-TH-F 0B-PLY-050-1220-2440-BT-TH-F
<!-- "6PMS_1RUN".run41protb --> ok 只有 0B-PLY-050-1220-2440-BT-TH-F
<!-- "6PMS_1RUN".run51instocktb -->0B-PLY-050-1220-2440-BT-TH-F
<!-- "6PMS_2OUT".out11hubtb -->0B-PLY-050-1220-2440-BT-TH-F
<!-- "6PMS_2OUT".out21rtntb  --> 沒東西
<!-- "2WMS_2OUT".out11paytb --> 0B-PLY-050-1220-2440-BT-TH-F
<!-- "2WMS_1RUN".run21txntb --> 0B-PLY-050-1220-2440-BT-TH-F
<!-- "2WMS_1RUN".run01opentb --> 沒有
"2WMS_1RUN".run10inv


out21(空的)
01_CKES_192.168.50.20   192.168.50.20 55432

-- 1. 更新第一个值
UPDATE "0ADM_2BDM".bdm00pdt
SET PN = '0B-PLY-030-1220-2440-BT-TH-F'
WHERE PN = '0B-PLY-030-1220-2440-HB-TH-F';

-- 2. 更新第二个值
UPDATE "0ADM_2BDM".bdm00pdt
SET PN = '0B-PLY-120-1220-2440-BT-TH-F'
WHERE PN = '0B-PLY-120-1220-2440-HB-TH-F';

-- 3. 更新第三个值
UPDATE "0ADM_2BDM".bdm00pdt
SET PN = '0B-PLY-150-1220-2440-BT-TH-F'
WHERE PN = '0B-PLY-150-1220-2440-HB-TH-F';

-- 4. 更新第四个值
UPDATE "0ADM_2BDM".bdm00pdt
SET PN = '0B-PLY-050-1220-2440-HB-TH-F'
WHERE PN = '0B-PLY-050-1220-2440-DH-TH-F';

-- 5. 更新第五个值
UPDATE "0ADM_2BDM".bdm00pdt
SET PN = '0B-PLY-050-1220-2440-BT-TH-F'
WHERE PN = '0B-PLY-050-1220-2440-HB-TH-F';
