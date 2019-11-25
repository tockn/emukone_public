
-- +migrate Up
START TRANSACTION ;
INSERT INTO users (id, name, identifier, icon_url, header_image_url, meta_description) VALUES(1, 'andymori', 'artist', 'http://andymori.com/img/top/201308280946_09166.jpg', 'header', '行こうぜベースマン');
INSERT INTO users (id, name, identifier, icon_url, header_image_url, meta_description) VALUES(2, 'artistuser2', 'artist', 'icon', 'header', 'meta');
INSERT INTO users (id, name, identifier, icon_url, header_image_url, meta_description) VALUES(3, 'ブッカー中野', 'booker', 'http://work-switch.persol-pt.co.jp/wp-content/uploads/2017/02/pierre_001.jpg', 'header', '立川を盛り上げる');
INSERT INTO users (id, name, identifier, icon_url, header_image_url, meta_description) VALUES(4, 'bookseruser2', 'booker', 'icon', 'header', 'meta');
INSERT INTO artist_users (user_id, why_description, how_description, free_description) VALUES (1, 'ベンガルトラの勇ましさと、ウイスキーの美味しさを布教するために活動しています', 'how1', 'andymoriは、2014年10月15日（水）日本武道館公演「andymori ラストライブ」をもって、
解散いたしました。

2007年から約7年間の活動を応援してくださったすべてのファンの皆さま、
支えてくださった関係者の皆さまに、心より御礼申し上げます。

andymoriは解散いたしますが、andymoriが生んだ数々の音楽は消えることはありません。
いつまでも皆さまの、そして未来の誰かの日々の中に生き続け、
その方の心と共にあり続ける音楽として愛していただけたら、この上なく幸せです。

本当にありがとうございました。');
INSERT INTO artist_users (user_id, why_description, how_description, free_description) VALUES (2, 'why2', 'how2', 'free2');
INSERT INTO booker_users (user_id, why_description, how_description, free_description) VALUES (3, '私の地元 ”立川” の次世代シ―ンを盛り上げたい！という気持ちで、全国のバンドを立川にお呼びしています。これまでに"赤い公園”の皆様にも出演して頂きました.', '主に立川を中心にイベントを立ち上げています！', '熱意のあるバンドをどんどん誘っていきます！');
INSERT INTO booker_users (user_id, why_description, how_description, free_description) VALUES (4, 'why2', 'how2', 'free2');
INSERT INTO user_tags (name) VALUES ('Rock');
INSERT INTO user_tags (name) VALUES ('西荻窪');
INSERT INTO user_tags (name) VALUES ('３ピース');
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (1, 1);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (1, 2);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (2, 2);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (2, 3);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (3, 1);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (3, 2);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (4, 2);
INSERT INTO user_tags_related_users (user_id, user_tag_id) VALUES (4, 3);
INSERT INTO locations (name) VALUES ('yokohama');
INSERT INTO locations (name) VALUES ('shibuya');
INSERT INTO locations (name) VALUES ('simokita');
INSERT INTO locations_related_users (user_id, location_id) VALUES (1, 1);
INSERT INTO locations_related_users (user_id, location_id) VALUES (2, 1);
INSERT INTO locations_related_users (user_id, location_id) VALUES (2, 2);
INSERT INTO locations_related_users (user_id, location_id) VALUES (3, 2);
INSERT INTO locations_related_users (user_id, location_id) VALUES (4, 2);
INSERT INTO locations_related_users (user_id, location_id) VALUES (4, 3);
INSERT INTO ichioshis (service, url, embed, title, artist_user_id) VALUES ('youtube', 'https://www.youtube.com/watch?v=-tg6WdgiDm8', '-tg6WdgiDm8', '革命', 1);
INSERT INTO ichioshis (service, url, embed, title, artist_user_id) VALUES ('youtube', 'https://youtube.com', '-tg6WdgiDm8', '革命', 2);
INSERT INTO ichioshis (service, url, embed, title, artist_user_id) VALUES ('eggs', 'https://eggs.mu/song/6e98644fcc8aa5620772cc749cd1ae61f517d8a7f9f7fe20b42ee78133d07e14', '', '天国', 2);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('taro', 'guitar', 1);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('sugawara', 'bass', 1);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('yamazaki', 'guitar', 2);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('sugawara', 'bass', 2);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('takuto', 'guitar&vo', 2);
INSERT INTO artist_members (name, part, artist_user_id) VALUES ('Fujishiro', 'dr', 2);
INSERT INTO user_images (image_url, user_id) VALUES ('https://image', 1);
INSERT INTO user_images (image_url, user_id) VALUES ('https://image', 2);
INSERT INTO user_images (image_url, user_id) VALUES ('https://image', 3);
INSERT INTO user_images (image_url, user_id) VALUES ('https://image', 4);
INSERT INTO website_urls (type, url, user_id) VALUES ('hp', 'https://hq', 1);
INSERT INTO website_urls (type, url, user_id) VALUES ('hp', 'https://hq', 2);
INSERT INTO website_urls (type, url, user_id) VALUES ('hp', 'https://hq', 3);
INSERT INTO website_urls (type, url, user_id) VALUES ('hp', 'https://hq', 4);
COMMIT;

-- +migrate Down
