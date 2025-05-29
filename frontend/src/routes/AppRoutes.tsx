import { Routes, Route, Navigate } from 'react-router-dom'
import { AuthProvider } from '../auth/AuthContext'
import ProtectedRoute from '../auth/ProtectedRoute'
import LoginPage from '../pages/LoginPage'
import DashboardLayout from '../layout/DashboardLayout'
import DashboardHome from '../pages/DashboardHome'

export default function AppRoutes() {
  return (
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Navigate to="/login" />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/dashboard" element={
          <ProtectedRoute>
            <DashboardLayout />
          </ProtectedRoute>
        }>
          <Route index element={<DashboardHome />} />
        </Route>
      </Routes>
    </AuthProvider>
  )
}
