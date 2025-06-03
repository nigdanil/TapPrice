// components/CategoryList.tsx
import { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

type Category = {
  id: number;
  name: string;
};

const CategoryList = () => {
  const [categories, setCategories] = useState<Category[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    axios.get('/api/categories')
      .then(res => {
        console.log('DEBUG categories response:', res.data);
        // Предполагаем, что категории находятся в res.data.categories
        const data = Array.isArray(res.data) ? res.data : res.data.categories;
        setCategories(data || []);
      })
      .catch(err => console.error('Ошибка загрузки категорий', err));
  }, []);

  return (
    <div className="grid grid-cols-2 gap-4 p-4">
      {categories.map((cat) => (
        <div
          key={cat.id}
          className="border p-4 rounded-xl shadow cursor-pointer hover:bg-gray-100"
          onClick={() => navigate(`/gallery/category/${cat.id}`)}
        >
          <h2 className="text-xl font-bold">{cat.name}</h2>
        </div>
      ))}
    </div>
  );
};

export default CategoryList;
