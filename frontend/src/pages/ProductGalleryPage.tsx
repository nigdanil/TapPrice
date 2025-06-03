// pages/ProductGalleryPage.tsx
import { Route, Routes } from 'react-router-dom';
import CategoryList from '../components/CategoryList';
import ProductGrid from '../components/ProductGrid';
import ProductDetail from '../components/ProductDetail';

const ProductGalleryPage = () => {
  return (
    <Routes>
      <Route path="/" element={<CategoryList />} />
      <Route path="category/:id" element={<ProductGrid />} />
      <Route path="product/:id" element={<ProductDetail />} />
    </Routes>
  );
};

export default ProductGalleryPage;