CREATE TABLE SamplingLetter (
    id SERIAL PRIMARY KEY ,
    letter VARCHAR(1) NOT NULL
);


INSERT INTO SamplingLetter (letter) VALUES 
    ('A'), ('B'), ('C'), ('D'), ('E'), ('F'), ('G'), ('H'), ('I'), ('J'), ('K'), ('L'), ('M'), ('N'), ('O'), ('P'), ('Q');

--

CREATE TABLE inspection_level (
    id SERIAL PRIMARY KEY,
    level VARCHAR(10) NOT NULL
);

CREATE TABLE product_qty_range (
    id SERIAL PRIMARY KEY,
    qty_from INT NOT NULL,
    qty_to INT NOT NULL
);

-- 插入 product_qty_range
ALTER TABLE product_qty_range
DROP COLUMN qty_to;

-- 新增一個可為 NULL 的 qty_to 欄位
ALTER TABLE product_qty_range
ADD COLUMN qty_to INT NULL;
INSERT INTO product_qty_range (qty_from, qty_to) VALUES
    (2, 8),
    (9, 15),
    (16, 25),
    (26, 50),
    (51, 90),
    (91, 150),
    (151, 280),
    (281, 500),
    (501, 1200),
    (1201, 3200),
    (3201, 10000),
    (10001, 35000),
    (35001, 150000),
    (150001, 500000),
    (500001, NULL); -
--

CREATE TABLE level_range_letter (
    id SERIAL PRIMARY KEY,
    inspection_level_id INT REFERENCES inspection_level(id),
    product_qty_range_id INT REFERENCES product_qty_range(id),
    sampling_letter_id INT REFERENCES Sampling_letter(id)
);

--- 新增對應level range letter
-- 插入 qty_from: 9 to:15 對應 level: I letter: A
-- INSERT INTO level_range_letter (inspection_level_id, product_qty_range_id, sampling_letter_id)
-- VALUES (
--     (SELECT id FROM inspection_level WHERE level = 'I'),
--     (SELECT id FROM product_qty_range WHERE qty_from = 9 AND qty_to = 15),
--     (SELECT id FROM Sampling_letter WHERE letter = 'A')
-- );

-- -- 插入 qty_from: 16 to:25 對應 level: I letter: B
-- INSERT INTO level_range_letter (inspection_level_id, product_qty_range_id, sampling_letter_id)
-- VALUES (
--     (SELECT id FROM inspection_level WHERE level = 'I'),
--     (SELECT id FROM product_qty_range WHERE qty_from = 16 AND qty_to = 25),
--     (SELECT id FROM Sampling_letter WHERE letter = 'B')
-- );

-- -- 插入 qty_from: 3201 to:10000 對應 level: II letter: L
-- INSERT INTO level_range_letter (inspection_level_id, product_qty_range_id, sampling_letter_id)
-- VALUES (
--     (SELECT id FROM inspection_level WHERE level = 'II'),
--     (SELECT id FROM product_qty_range WHERE qty_from = 3201 AND qty_to = 10000),
--     (SELECT id FROM Sampling_letter WHERE letter = 'L')
-- );


-- 給gpt的資料
qty_from: 9  
to:15
level : I
letter A
----
qty_from: 16
to:25
level : I
letter B
-----
qty_from: 3201  
to:10000
level : II
letter L
----
qty_from: 16
to:25
level : II
letter C

---
INSERT INTO level_range_letter (inspection_level, qty_from, qty_to, sampling_letter)
VALUES
    ('I', 9, 15, 'A'),
    ('I', 16, 25, 'B'),
    ('II', 3201, 10000, 'L'),
    ('II', 16, 25, 'C'),
    ('II', 1201, 3200, 'K');


-- 1. 插入多一個letter R
INSERT INTO sampling_letter (letter) VALUES ('R');

-- 2. 幫我加上一column叫做 sample_size
ALTER TABLE sampling_letter ADD COLUMN sample_size INT;

-- 更新 sample_size 資料
UPDATE sampling_letter SET sample_size = CASE
    WHEN id = 1 THEN 2
    WHEN id = 2 THEN 3
    WHEN id = 3 THEN 5
    WHEN id = 4 THEN 8
    WHEN id = 5 THEN 13
    WHEN id = 6 THEN 20
    WHEN id = 7 THEN 32
    WHEN id = 8 THEN 50
    WHEN id = 9 THEN 80
    WHEN id = 10 THEN 125
    WHEN id = 11 THEN 200
    WHEN id = 12 THEN 315
    WHEN id = 13 THEN 500
    WHEN id = 14 THEN 800
    WHEN id = 15 THEN 1250
    WHEN id = 16 THEN 2000
END;

--
DELETE FROM Sampling_letter WHERE letter IN ('I', 'O');

-- 更新值得判斷依據
UPDATE Sampling_letter
SET sample_size = CASE
    WHEN letter = 'A' THEN 2
    WHEN letter = 'B' THEN 3
    WHEN letter = 'C' THEN 5
    WHEN letter = 'D' THEN 8
    WHEN letter = 'E' THEN 13
    WHEN letter = 'F' THEN 20
    WHEN letter = 'G' THEN 32
    WHEN letter = 'H' THEN 50
    WHEN letter = 'J' THEN 80
    WHEN letter = 'K' THEN 125
    WHEN letter = 'L' THEN 200
    WHEN letter = 'M' THEN 315
    WHEN letter = 'N' THEN 500
    WHEN letter = 'P' THEN 800
    WHEN letter = 'Q' THEN 1250
    WHEN letter = 'R' THEN 2000
END;

--
CREATE TABLE aql (
    id SERIAL PRIMARY KEY,
    aql DECIMAL(4, 3) NOT NULL
);

CREATE TABLE single_plan (
    id SERIAL PRIMARY KEY,
    aql int,
    sampling_letter_id INT NOT NULL,
    ac_num INT NOT NULL,
    re_num INT NOT NULL,
    FOREIGN KEY (sampling_letter_id) REFERENCES sampling_letter(id),
	foreign key (aql) REFERENCES aql(id)
);

--
-- 插入 aql 表的資料
INSERT INTO aql (aql) VALUES
(0.065),
(0.10),
(0.15),
(0.25),
(0.40),
(0.65),
(1.0),
(1.5),
(2.5),
(4.0),
(6.5);
---
CREATE TABLE user_sampling_plan (
    id SERIAL PRIMARY KEY,
    user_id INT,

    product_qty INT,

    inspection_level char(1),
    sampling_letter char(1),

    critical_aql_id double precision,
    major_aql double precision,
    minor_aql double precision,
    
    qty_range_from int,
    qty_range_to int,
    
    critical_ac INT,
    major_ac INT,
    minor_ac INT,
    
    -- critical_re INT,
    -- major_re INT,
    -- minor_re INT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE single_plan
ALTER COLUMN sampling_letter TYPE INT;

ALTER TABLE single_plan
RENAME COLUMN sampling_letter_id TO sampling_letter;

aql: 2.5
letter: L
ac:10
re:11
---
aql:4.0
letter:L
ac:14
re:15
---
aql:0.065
letter:L
ac:0
re:1