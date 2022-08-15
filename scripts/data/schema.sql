DROP TABLE IF EXISTS users_tab;
CREATE TABLE users_tab (

	user_id BIGINT NOT NULL AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	create_time INT NOT NULL,
    update_time INT NOT NULL,

    PRIMARY KEY (user_id)
);
CREATE UNIQUE INDEX idx_user_id ON users_tab(user_id);

DROP TABLE IF EXISTS products_tab;
CREATE TABLE products_tab (
    
    product_id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    URL VARCHAR(255) NOT NULL,
    view_number INT NOT NULL DEFAULT 0,
	create_time INT NOT NULL,
    update_time INT NOT NULL,

    PRIMARY KEY (product_id)
);
CREATE UNIQUE INDEX idx_product_id ON products_tab(product_id);

DROP TABLE IF EXISTS categories_tab;
CREATE TABLE categories_tab (
    category_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
	create_time INT NOT NULL,
    update_time INT NOT NULL,

    PRIMARY KEY(category_id)
);
CREATE UNIQUE INDEX idx_category_id ON categories_tab(user_id);

DROP TABLE IF EXISTS product_category_tab;
CREATE TABLE product_category_tab (
    product_category_id BIGINT NOT NULL AUTO_INCREMENT,
    product_id BIGINT NOT NULL,
    category_id INT NOT NULL,

    PRIMARY KEY (product_category_id)
);
CREATE UNIQUE INDEX idx_product_category_id ON product_category_tab(product_category_id);

DROP TABLE IF EXISTS comments_tab;
CREATE TABLE comments_tab (
    comment_id BIGINT NOT NULL AUTO_INCREMENT,
    text VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
	create_time INT NOT NULL,
    update_time INT NOT NULL,

    PRIMARY KEY (comment_id)
);
CREATE UNIQUE INDEX idx_comment_id ON comments_tab(comment_id);

DROP TABLE IF EXISTS replies_tab;
CREATE TABLE replies_tab (
    id BIGINT NOT NULL AUTO_INCREMENT,
    parent_id BIGINT NOT NULL,
    reply_id BIGINT NOT NULL,

    PRIMARY KEY (id)
);
