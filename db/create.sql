CREATE DATABASE `TEAMC` DEFAULT CHARACTER SET utf8mb4;
USE `TEAMC`;
DROP TABLE IF EXISTS shops;

CREATE TABLE IF NOT EXISTS shops (
  id integer PRIMARY KEY AUTO_INCREMENT,
  shop_name varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  1,
  "栗"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  2,
  "おんどり"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  3,
  "La Terrasse “irisée”"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  4,
  "カフェ エトランジェ ナラッド"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  5,
  "ALL DAY DINING"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  6,
  "中川政七商店　奈良の工芸に触れる体験"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  7,
  "今西清兵衛商店"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  8,
  "ならまち格子の家"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  9,
  "瑜伽山園地"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  10,
  "寧楽美術館"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  11,
  "日本酒とおつまみ　chuin"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  12,
  "なら泉勇斎"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  13,
  "樫舎"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  14,
  "SUNNY and MORE"
);

INSERT INTO shops (
  id,
  shop_name
)
VALUES (
  15,
  "ボリクコーヒー"
);


-- "昼食、夕食	栗	34.68096923516614, 135.83505440564
-- 昼食、夕食	おんどり	34.697889069751675, 135.8452012810097
-- 昼食、夕食	La Terrasse “irisée” 	34.70355272581798, 135.76559570736123
-- 昼食、夕食	カフェ エトランジェ ナラッド 	34.682309921526944, 135.82655925771624
-- 昼食、夕食	ALL DAY DINING	34.683420861642915, 135.82780645642984
-- 30分以上	中川政七商店　奈良の工芸に触れる体験	34.66844085304311, 135.83071231303978
-- 30分以上	今西清兵衛商店	34.6771167799442, 135.8346802547372
-- 30分以上	ならまち格子の家	34.67519626978004, 135.83075745458194
-- 30分以上	瑜伽山園地	34.6795394700259, 135.83924008867956
-- 30分以上	寧楽美術館	34.68621397298806, 135.83719518341763
-- 30分以下	日本酒とおつまみ　chuin	34.67879867636942, 135.83066651225275
-- 30分以下	なら泉勇斎	34.679852836545734, 135.82983086566745
-- 30分以下	樫舎	34.67875598474761, 135.8312995000453
-- 30分以下	SUNNY and MORE	34.679749291117396, 135.8312133815698
-- 30分以下	ボリクコーヒー	34.67665865074141, 135.83004641964288"