// components/ProductDetail.tsx
import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';

const ProductDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [product, setProduct] = useState<any>(null);

  useEffect(() => {
    axios.get(`/api/product/${id}`)
      .then(res => setProduct(res.data))
      .catch(err => console.error('Ошибка загрузки продукта', err));
  }, [id]);

  if (!product) return <div className="p-4">Загрузка...</div>;

  return (
    <div className="p-4">
      <button onClick={() => navigate(-1)} className="mb-4 text-blue-500 hover:underline">
        ← Назад
      </button>

      <h1 className="text-2xl font-bold mb-2">{product.name}</h1>

      {product.description && (
        <p className="mb-2 text-gray-700"><strong>Описание:</strong> {product.description}</p>
      )}

      {product.composition && (
        <p className="mb-2 text-gray-700"><strong>Состав:</strong> {product.composition}</p>
      )}

      {Array.isArray(product.cert_links) && product.cert_links.length > 0 && (
        <div className="mt-4">
          <strong>Сертификаты:</strong>
          <ul className="list-disc list-inside text-blue-600">
            {product.cert_links.map((link: string, index: number) => (
              <li key={index}>
                <a href={link} target="_blank" rel="noopener noreferrer" className="underline">
                  Скачать сертификат {index + 1}
                </a>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default ProductDetail;
