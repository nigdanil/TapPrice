import { useEffect, useState } from "react"

type Category = {
  id: number
  name: string
}

type Product = {
  id: number
  name: string
  description: string
  category: string
  venue: string
}

export default function MenuPage() {
  const [categories, setCategories] = useState<Category[]>([])
  const [products, setProducts] = useState<Product[]>([])
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null)
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null)

  useEffect(() => {
    fetch("/api/categories")
      .then((res) => res.json())
      .then(setCategories)
      .catch((err) => console.error("Ошибка загрузки категорий:", err))
  }, [])

  useEffect(() => {
    if (selectedCategory !== null) {
      fetch(`/api/products?category_id=${selectedCategory}`)
        .then((res) => res.json())
        .then(setProducts)
        .catch((err) => console.error("Ошибка загрузки продуктов:", err))
    }
  }, [selectedCategory])

  if (selectedProduct) {
    return (
      <div style={{ padding: "1rem" }}>
        <button onClick={() => setSelectedProduct(null)}>← Назад к продуктам</button>
        <h2>{selectedProduct.name}</h2>
        <p>{selectedProduct.description}</p>
        <p>Категория: {selectedProduct.category}</p>
        <p>Площадка: {selectedProduct.venue}</p>
      </div>
    )
  }

  return (
    <div style={{ padding: "1rem" }}>
      <h1>Меню</h1>

      {!selectedCategory ? (
        <>
          <h2>Категории</h2>
          <div style={{ display: "flex", flexWrap: "wrap", gap: "1rem" }}>
            {categories.map((cat) => (
              <div
                key={cat.id}
                style={{
                  border: "1px solid #ccc",
                  padding: "1rem",
                  cursor: "pointer",
                  borderRadius: "8px",
                }}
                onClick={() => setSelectedCategory(cat.id)}
              >
                {cat.name}
              </div>
            ))}
          </div>
        </>
      ) : (
        <>
          <button onClick={() => setSelectedCategory(null)}>← Назад к категориям</button>
          <h2>Продукты</h2>
          <div style={{ display: "flex", flexWrap: "wrap", gap: "1rem" }}>
            {products.map((prod) => (
              <div
                key={prod.id}
                style={{
                  border: "1px solid #ccc",
                  padding: "1rem",
                  width: "250px",
                  borderRadius: "8px",
                  cursor: "pointer",
                }}
                onClick={() => setSelectedProduct(prod)}
              >
                <strong>{prod.name}</strong>
                <p>{prod.description}</p>
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  )
}
