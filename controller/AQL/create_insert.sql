CREATE TABLE "aql".user_sampling_plan (
	    id SERIAL PRIMARY KEY,
	    user_id INT,
	
	    product_qty INT,
	
	    inspection_level char(1),
	    sampling_letter char(1),
	
	    critical_aql double precision,
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
	
	CREATE TABLE "aql".single_plan (
	    id SERIAL PRIMARY KEY,
	    ac_num INT NOT NULL,
	    re_num INT NOT NULL,
	    sampling_letter char(1),
	    aql double precision
	);
	
	CREATE TABLE "aql".sampling_letter (
	    id SERIAL PRIMARY KEY,
	    letter VARCHAR(1) NOT NULL,
	    size int NOT NULL
	);
	
	
	CREATE TABLE "aql".level_range_letter (
	    id SERIAL PRIMARY KEY,
	    inspection_level VARCHAR(5) NOT NULL,
	    qty_range_from INT NOT NULL,
	    qty_range_to INT NOT NULL,
	    sampling_letter CHAR(1) not null
	);
	
	
	CREATE TABLE "aql".product_qty_range (
	    id SERIAL PRIMARY KEY,
	    qty_from INT NOT NULL,
	    qty_to INT NOT NULL
	);
	
	CREATE TABLE "aql".inspection_level (
	    id SERIAL PRIMARY KEY,
	    level VARCHAR(10) NOT NULL
	);
	
	
	CREATE TABLE "aql".aql (
	    id SERIAL PRIMARY KEY,
	    aql DECIMAL(4, 3) NOT NULL
	);
	
----- insert
-- ALTER TABLE "aql".product_qty_range
-- ALTER COLUMN qty_to DROP NOT NULL;

INSERT INTO "aql".single_plan (aql, sampling_letter, ac_num, re_num)
VALUES
    (2.5, 'L', 10, 11),
    (4.0, 'L', 14, 15),
    (0.065, 'L', 0, 1),
	 (0.65, 'L', 3, 4)

INSERT INTO "aql".inspection_level (level)
VALUES
    ('I'),
    ('II'),
    ('III'),
    ('S-1'),
    ('S-2'),
    ('S-3');

INSERT INTO "aql".aql (aql) VALUES
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

INSERT INTO "aql".sampling_letter (letter,size) VALUES 
    ('A',2), ('B',3), ('C',5), ('D',8), ('E',13),
	('F',20), ('G',32), ('H',50), ('J',80), 
	('K',125), ('L',200), ('M',315), ('N',500), 
	('P',800), ('Q',1250),('R',2000);
	


INSERT INTO "aql".product_qty_range (qty_from, qty_to) VALUES
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
    (500001, NULL);
	
	
	INSERT INTO "aql".level_range_letter (inspection_level, qty_range_from, qty_range_to, sampling_letter)
VALUES
    ('I', 9, 15, 'A'),
    ('I', 16, 25, 'B'),
    ('II', 3201, 10000, 'L'),
    ('II', 16, 25, 'C'),
    ('II', 1201, 3200, 'K');