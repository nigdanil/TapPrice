// D:\TapPrice\frontend\src\auth\ProtectedRoute.tsx
import { Navigate } from 'react-router-dom'
import { useAuth } from './AuthContext'
import type { ReactNode } from 'react'

type ProtectedRouteProps = {
  children: ReactNode
}

export default function ProtectedRoute({ children }: ProtectedRouteProps) {
  const { user, loading } = useAuth()

  if (loading) return <div>Загрузка...</div>
  if (!user) return <Navigate to="/login" />

  return <>{children}</>
}
