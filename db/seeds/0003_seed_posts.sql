INSERT INTO posts (id, user_id, caption, photo)
VALUES ('10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003', 'Foto pertama saya di pegunungan, udara segar!', 'photo_charlie_1.jpg')
ON CONFLICT DO NOTHING;

INSERT INTO posts (id, user_id, caption, photo)
VALUES ('10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000003', 'Sunset di pantai, hasil bidikan terbaik.', 'photo_charlie_2.jpg')
ON CONFLICT DO NOTHING;

INSERT INTO posts (id, user_id, caption, photo)
VALUES ('10000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000002', 'Menulis draft blog baru tentang kota tua.', 'photo_bob_1.jpg')
ON CONFLICT DO NOTHING;