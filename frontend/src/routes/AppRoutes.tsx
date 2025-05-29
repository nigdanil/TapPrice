import { Routes, Route, Navigate } from 'react-router-dom'
import { AuthProvider } from '../auth/AuthContext'
import ProtectedRoute from '../auth/ProtectedRoute'

// Страницы
import LoginPage from '../pages/LoginPage'
import ProductsPage from '../pages/ProductsPage'
import MenuPage from '../pages/MenuPage'

// Панель администратора
import DashboardLayout from '../layout/DashboardLayout'
import DashboardHome from '../pages/DashboardHome'

export default function AppRoutes() {
  return (
    <AuthProvider>
      <Routes>
        {/* Публичные маршруты */}
        <Route path="/" element={<Navigate to="/menu" />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/menu" element={<MenuPage />} />
        <Route path="/products" element={<ProductsPage />} />

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
