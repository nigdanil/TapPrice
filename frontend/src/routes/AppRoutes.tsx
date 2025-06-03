import { Routes, Route, Navigate } from 'react-router-dom'
import { AuthProvider } from '../auth/AuthContext'
import ProtectedRoute from '../auth/ProtectedRoute'

// Страницы
import LoginPage from '../pages/LoginPage'
import ProductsPage from '../pages/ProductsPage'
import ProductGalleryPage from '../pages/ProductGalleryPage' // ← добавлено

// Панель администратора
import DashboardLayout from '../layout/DashboardLayout'
import DashboardHome from '../pages/DashboardHome'

export default function AppRoutes() {
  return (
    <AuthProvider>
      <Routes>
        {/* Публичные маршруты */}
        <Route path="/" element={<Navigate to="/gallery" />} /> {/* ← теперь ведёт в /gallery */}
        <Route path="/login" element={<LoginPage />} />
        <Route path="/products" element={<ProductsPage />} />
        <Route path="/gallery/*" element={<ProductGalleryPage />} /> {/* ← добавлен маршрут */}

        {/* Защищённые маршруты с layout и outlet */}
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <DashboardLayout />
            </ProtectedRoute>
          }
        >
          <Route index element={<DashboardHome />} />
          {/* Здесь можно добавить другие вложенные маршруты */}
        </Route>
      </Routes>
    </AuthProvider>
  )
}
