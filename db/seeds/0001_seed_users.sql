
INSERT INTO users (id, fullname, email, password, bio)
VALUES ('00000000-0000-0000-0000-000000000001', 'Alice Admin', 'alice@mail.com', 'HASHED_PASSWORD_ALICE', 'Akun admin, menguji segalanya.')
ON CONFLICT DO NOTHING;
INSERT INTO users (id, fullname, email, password, bio)
VALUES ('00000000-0000-0000-0000-000000000002', 'Bob Blogger', 'bob@mail.com', 'HASHED_PASSWORD_BOB', 'Hobi nulis dan travelling.')
ON CONFLICT DO NOTHING;

INSERT INTO users (id, fullname, email, password, bio)
VALUES ('00000000-0000-0000-0000-000000000003', 'Charlie Creator', 'charlie@mail.com', 'HASHED_PASSWORD_CHARLIE', 'Fokus pada fotografi alam.')
ON CONFLICT DO NOTHING;

INSERT INTO users (id, fullname, email, password, bio)
VALUES ('00000000-0000-0000-0000-000000000004', 'Diana Dummy', 'diana@mail.com', 'HASHED_PASSWORD_DIANA', 'Hanya pengamat.')
ON CONFLICT DO NOTHING;