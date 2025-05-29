#!/bin/bash

docker exec -i tapprice-postgres-1 psql -U user -d digital_labels <<'SQL'
-- Добавим venue и сохраним его id
WITH inserted_venue AS (
  INSERT INTO venues (name, slug) VALUES ('Demo Venue', 'demo')
  ON CONFLICT (slug) DO NOTHING
  RETURNING id
), inserted_category AS (
  -- Используем id из inserted_venue
  INSERT INTO categories (venue_id, name)
  SELECT id, 'Beer' FROM inserted_venue
  RETURNING id, venue_id
)
-- Финально создаём продукт
INSERT INTO products (venue_id, category_id, name, description, composition, cert_links)
SELECT venue_id, id, 'Demo Beer', 'Great taste', 'water, hops, malt',
  ARRAY['https://example.com/cert.pdf']
FROM inserted_category;
SQL
