create database rent_amera_db;

-- USERS TABLE
CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    deposit_amount NUMERIC(15,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- END OF USERS

-- CAMERAS TABLE
CREATE TABLE cameras (
    camera_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    rental_cost NUMERIC(15,2) NOT NULL,
    category VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- END OF CAMERAS

-- RENTAL ORDERS
CREATE TABLE rental_orders (
    rental_order_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    total_amount NUMERIC(15,2) NOT NULL DEFAULT 0,
    status VARCHAR(30) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE rental_orders
ADD CONSTRAINT fk_rental_order_user
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE;
-- END OF RENTAL ORDERS


CREATE TABLE rental_order_details (
    rental_detail_id BIGSERIAL PRIMARY KEY,
    rental_order_id BIGINT NOT NULL,
    camera_id BIGINT NOT NULL,
    rental_cost NUMERIC(15,2) NOT NULL,
    rental_days INT NOT NULL DEFAULT 1,
    subtotal NUMERIC(15,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- 
ALTER TABLE rental_order_details
ADD CONSTRAINT fk_rental_detail_order
FOREIGN KEY (rental_order_id)
REFERENCES rental_orders(rental_order_id)
ON DELETE CASCADE;

ALTER TABLE rental_order_details
ADD CONSTRAINT fk_rental_detail_camera
FOREIGN KEY (camera_id)
REFERENCES cameras(camera_id)
ON DELETE RESTRICT;


CREATE TABLE payments (
    payment_id BIGSERIAL PRIMARY KEY,
    rental_order_id BIGINT NOT NULL UNIQUE,
    amount NUMERIC(15,2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    payment_status VARCHAR(30) NOT NULL DEFAULT 'PENDING',
    payment_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE payments 
ADD CONSTRAINT fk_payment_rental_order
FOREIGN KEY (rental_order_id)
REFERENCES rental_orders(rental_order_id)
ON DELETE CASCADE;

ALTER TABLE payments 
ADD CONSTRAINT fk_payment_verified_by
FOREIGN KEY (verified_by)
REFERENCES users(user_id)
ON DELETE SET NULL;

CREATE TABLE top_ups (
    top_up_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE top_ups
ADD CONSTRAINT fk_topups_user
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE;

ALTER TABLE top_ups
ADD CONSTRAINT check_topup_amount
CHECK (amount > 0);


select * from users u ;
select * from top_ups tu ;
