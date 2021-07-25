CREATE DATABASE `TEAMC` DEFAULT CHARACTER SET utf8mb4;
USE `TEAMC`;
DROP TABLE IF EXISTS shops;

CREATE TABLE IF NOT EXISTS shops (
  id integer PRIMARY KEY AUTO_INCREMENT,
  shop_name varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  self_intro varchar(100),
  category_1_id integer,
  category_2_id integer,
  category_3_id integer,
  latitude varchar(100),
  longitude varchar(100)
);

DROP TABLE IF EXISTS recommends;

CREATE TABLE IF NOT EXISTS recommends (
  id integer PRIMARY KEY AUTO_INCREMENT,
  recommender_id integer,
  recommended_shop_id integer,
  intro varchar(100)
);

DROP TABLE IF EXISTS tourist_comments;

CREATE TABLE IF NOT EXISTS tourist_comments (
  id integer PRIMARY KEY AUTO_INCREMENT,
  shop_id integer,
  commenter varchar(100),
  comment varchar(100)
);

DROP TABLE IF EXISTS categories;

CREATE TABLE IF NOT EXISTS categories (
  id integer PRIMARY KEY AUTO_INCREMENT,
  category_name varchar(100)
);

INSERT INTO categories (category_name) VALUES ("ランチ");
INSERT INTO categories (category_name) VALUES ("ディナー");
INSERT INTO categories (category_name) VALUES ("伝統工芸品");
INSERT INTO categories (category_name) VALUES ("ファミリー向け");
INSERT INTO categories (category_name) VALUES ("お一人様");
INSERT INTO categories (category_name) VALUES ("カップル向け");
INSERT INTO categories (category_name) VALUES ("歴史");
INSERT INTO categories (category_name) VALUES ("郷土料理");
INSERT INTO categories (category_name) VALUES ("隠れ家カフェ");
INSERT INTO categories (category_name) VALUES ("お土産");
INSERT INTO categories (category_name) VALUES ("若者に人気");
INSERT INTO categories (category_name) VALUES ("中高年に人気");

INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (1, "20代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (1, "30代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (1, "60代女性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (2, "70代男性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (2, "20代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (2, "30代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (3, "70代男性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (3, "40代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (3, "60代女性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (4, "20代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (4, "30代男性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (4, "70代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (5, "60代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (5, "20代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (5, "40代女性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (6, "70代男性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (6, "60代女性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (6, "20代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (7, "30代男性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (7, "40代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (7, "70代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (8, "30代男性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (8, "40代女性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (8, "60代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (9, "20代女性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (9, "40代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (9, "70代男性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (10, "60代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (10, "20代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (10, "40代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (11, "30代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (11, "70代男性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (11, "40代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (12, "60代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (12, "30代男性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (12, "70代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (13, "20代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (13, "30代男性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (13, "40代女性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (14, "60代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (14, "70代男性", "落ち着いた雰囲気でした。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (14, "20代女性", "貴重な体験でした！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (15, "30代男性", "最高！！！");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (15, "60代女性", "かなり良いです。");
INSERT INTO tourist_comments (shop_id, commenter, comment) VALUES (15, "20代女性", "落ち着いた雰囲気でした。");



INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  1,
  2,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  1,
  3,
  "いい感じです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  1,
  4,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  2,
  5,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  2,
  7,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  2,
  10,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  3,
  1,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  3,
  6,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  3,
  8,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  4,
  2,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  4,
  3,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  4,
  15,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  5,
  4,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  5,
  6,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  5,
  7,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  6,
  10,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  6,
  11,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  6,
  14,
  "すごいです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  7,
  9,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  7,
  13,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  7,
  14,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  8,
  9,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  8,
  11,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  8,
  15,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  9,
  3,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  9,
  5,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  9,
  10,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  10,
  2,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  10,
  6,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  10,
  12,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  11,
  1,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  11,
  4,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  11,
  7,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  12,
  9,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  12,
  13,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  12,
  14,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  13,
  2,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  13,
  7,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  13,
  9,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  14,
  10,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  14,
  11,
  "ぜひ行ってみて！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  14,
  12,
  "結構いいです。"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  15,
  4,
  "大好きです！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  15,
  12,
  "毎日通ってます！"
);

INSERT INTO recommends (
  recommender_id,
  recommended_shop_id,
  intro
)
VALUES (
  15,
  13,
  "結構いいです。"
);

-- 昼食、夕食	栗	34.68096923516614, 135.83505440564
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  1,
  "栗",
  "これはテストですがとても美味しいです。",
  4,
  5,
  11,
  "34.68096923516614",
  "135.83505440564"
);

-- 昼食、夕食	おんどり	34.697889069751675, 135.8452012810097
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  2,
  "おんどり",
  "これはテストですがとてもオススメです。",
  3,
  6,
  9,
  "34.697889069751675",
  "135.8452012810097"
);

-- 昼食、夕食	La Terrasse “irisée” 	34.70355272581798, 135.76559570736123
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  3,
  "La Terrasse “irisée”",
  "これはテストですがぜひ来てね！",
  1,
  5,
  9,
  "34.70355272581798",
  "135.76559570736123"
);

-- 昼食、夕食	カフェ エトランジェ ナラッド 	34.682309921526944, 135.82655925771624
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  4,
  "カフェ エトランジェ ナラッド",
  "これはテストですがたくさん来てくれるとうれしいです！",
  8,
  10,
  12,
  "34.682309921526944",
  "135.82655925771624"
);

-- 昼食、夕食	ALL DAY DINING	34.683420861642915, 135.82780645642984
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  5,
  "ALL DAY DINING",
  "これはテストですがたくさん食べてってね！",
  3,
  7,
  11,
  "34.683420861642915",
  "135.82780645642984"
);

-- 30分以上	中川政七商店　奈良の工芸に触れる体験	34.66844085304311, 135.83071231303978
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  6,
  "中川政七商店　奈良の工芸に触れる体験",
  "これはテストですが楽しんでいってね！",
  2,
  6,
  10,
  "34.66844085304311",
  "135.83071231303978"
);

-- 30分以上	今西清兵衛商店	34.6771167799442, 135.8346802547372
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  7,
  "今西清兵衛商店",
  "これはテストですがたくさん見てってください。",
  5,
  10,
  12,
  "34.6771167799442",
  "135.8346802547372"
);

-- 30分以上	ならまち格子の家	34.67519626978004, 135.83075745458194
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  8,
  "ならまち格子の家",
  "これはテストですがいろいろ見ていってみてください！",
  4,
  7,
  10,
  "34.67519626978004",
  "135.83075745458194"
);

-- 30分以上	瑜伽山園地	34.6795394700259, 135.83924008867956
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  9,
  "瑜伽山園地",
  "これはテストですが中を見回ってみてね！",
  6,
  8,
  10,
  "34.6795394700259",
  "135.83924008867956"
);

-- 30分以上	寧楽美術館	34.68621397298806, 135.83719518341763
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  10,
  "寧楽美術館",
  "これはテストですがたくさん美術品みてってね！",
  1,
  4,
  7,
  "34.68621397298806",
  "135.83719518341763"
);

-- 30分以下	日本酒とおつまみ　chuin	34.67879867636942, 135.83066651225275
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  11,
  "日本酒とおつまみ　chuin",
  "これはテストですが美味しいものいっぱいあるよ！",
  2,
  5,
  7,
  "34.67879867636942",
  "135.83066651225275"
);

-- 30分以下	なら泉勇斎	34.679852836545734, 135.82983086566745
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  12,
  "なら泉勇斎",
  "これはテストですがたくさん見てみてね！",
  4,
  5,
  6,
  "34.679852836545734",
  "135.82983086566745"
);

-- 30分以下	樫舎	34.67875598474761, 135.8312995000453
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  13,
  "樫舎",
  "これはテストですが楽しんでね！",
  10,
  11,
  12,
  "34.67875598474761",
  "135.8312995000453"
);

-- 30分以下	SUNNY and MORE	34.679749291117396, 135.8312133815698
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  14,
  "SUNNY and MORE",
  "これはテストですが面白いところですよ！",
  5,
  7,
  9,
  "34.679749291117396",
  "135.8312133815698"
);

-- 30分以下	ボリクコーヒー	34.67665865074141, 135.83004641964288"
INSERT INTO shops (
  id,
  shop_name,
  self_intro,
  category_1_id,
  category_2_id,
  category_3_id,
  latitude,
  longitude
)
VALUES (
  15,
  "ボリクコーヒー",
  "これはテストですがうちのコーヒー飲んでって！",
  1,
  2,
  3,
  "34.67665865074141",
  "135.83004641964288"
);