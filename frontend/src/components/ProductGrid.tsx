// components/ProductGrid.tsx
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';
import ProductCard from './ProductCard';

interface Product {
  id: number;
  name: string;
}

const ProductGrid = () => {
  const { id } = useParams();
  const [products, setProducts] = useState<Product[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    axios.get(`/api/products?category_id=${id}`)
      .then(res => {
        console.log("DEBUG products:", res.data);
        setProducts(res.data || []);
      })
      .catch(err => console.error("Ошибка загрузки продуктов", err));
  }, [id]);

  return (
    <div className="p-4">
      <button
        onClick={() => navigate('/gallery')}
        className="mb-4 px-4 py-2 border rounded hover:bg-gray-200"
      >
        ← Назад к категориям
      </button>

      <div className="grid grid-cols-2 gap-4">
        {products.map(product => (
          <ProductCard
            key={product.id}
            product={product}
            onClick={() => navigate(`/gallery/product/${product.id}`)}
          />
        ))}
      </div>
    </div>
  );
};

export default ProductGrid;
