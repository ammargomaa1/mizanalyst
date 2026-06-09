-- Seed a default admin user (password: "password123" hashed with bcrypt)
INSERT INTO users (name, email, phone_number, password, created_at, updated_at)
VALUES (
    'Admin',
    'admin@mizanalyst.com',
    '+1234567890',
    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
    NOW(),
    NOW()
) ON CONFLICT (email) DO NOTHING;
