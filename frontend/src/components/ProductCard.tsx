// components/ProductCard.tsx
interface ProductCardProps {
  product: any;
  onClick: () => void;
}

const ProductCard = ({ product, onClick }: ProductCardProps) => {
  return (
    <div onClick={onClick} className="border p-4 rounded-xl shadow hover:bg-gray-100 cursor-pointer">
      <h3 className="text-lg font-semibold">{product.name}</h3>
    </div>
  );
};

export default ProductCard;