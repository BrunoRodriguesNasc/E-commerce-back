-- Criar tabela de produtos
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    rating DECIMAL(3,2) DEFAULT 0.0,
    size VARCHAR(255) DEFAULT 'M',
    reviews TEXT[] DEFAULT '{}',
    image VARCHAR(255) DEFAULT 'https://images.unsplash.com/photo-placeholder'
);

-- Inserir dados na tabela de produtos
INSERT INTO products (name, description, price, rating, size, reviews, image)
VALUES
('Casual T-Shirt', 'Soft cotton t-shirt, perfect for everyday use.', 19.99, 4.5, 'M', 
ARRAY['Great quality!', 'Very comfortable.'], 
'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab'),

('Slim Fit Jeans', 'Classic slim fit jeans, available in dark wash.', 49.99, 4.2, 'L', 
ARRAY['Perfect fit!', 'Love the color.'], 
'https://images.unsplash.com/photo-1542272604-787c3835535d'),

('Leather Jacket', 'Stylish leather jacket with a comfortable lining.', 129.99, 4.8, 'L', 
ARRAY['Looks amazing!', 'Perfect for winter.'], 
'https://images.unsplash.com/photo-1551028719-00167b16eac5'),

('Running Shoes', 'Lightweight running shoes with excellent grip.', 79.99, 4.6, '42', 
ARRAY['Super comfortable!', 'Great for long runs.'], 
'https://images.unsplash.com/photo-1542291026-7eec264c27ff'),

('Denim Jacket', 'Trendy denim jacket with button closure.', 69.99, 4.4, 'M', 
ARRAY['Nice design!', 'Fits well.'], 
'https://images.unsplash.com/photo-1523205771623-e0faa4d2813d'),

('Floral Summer Dress', 'Lightweight and breathable summer dress.', 34.99, 4.7, 'S', 
ARRAY['Beautiful design!', 'Perfect for summer.'], 
'https://images.unsplash.com/photo-1572804013309-59a88b7e92f1'),

('Hoodie Sweatshirt', 'Cozy hoodie with adjustable drawstrings.', 39.99, 4.5, 'XL', 
ARRAY['Warm and comfy.', 'Great for cold days.'], 
'https://images.unsplash.com/photo-1556821840-3a63f95609a7'),

('Chino Pants', 'Slim-fit chino pants in a variety of colors.', 44.99, 4.3, 'M', 
ARRAY['Perfect for work!', 'Very comfortable fit.'], 
'https://images.unsplash.com/photo-1473966968600-fa801b869a1a'),

('Long Sleeve Shirt', 'Formal long sleeve shirt, perfect for office wear.', 29.99, 4.2, 'L', 
ARRAY['Elegant design.', 'Good material.'], 
'https://images.unsplash.com/photo-1596755094514-f87e34085b2c'),

('Ankle Boots', 'Durable ankle boots with modern design.', 89.99, 4.6, '41', 
ARRAY['Great quality!', 'Stylish and comfortable.'], 
'https://images.unsplash.com/photo-1542840843-3349799cded6');
