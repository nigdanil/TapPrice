// D:\TapPrice\frontend\src\auth\AuthContext.tsx
import { createContext, useContext, useEffect, useState } from 'react'
import Cookies from 'js-cookie'

const AuthContext = createContext<any>(null)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<boolean | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const token = Cookies.get('access_token')
    setUser(token ? true : null)
    setLoading(false)
  }, [])

  const login = () => {
    setUser(true)
  }

  const logout = () => {
    Cookies.remove('access_token')
    setUser(null)
  }

  return (
    <AuthContext.Provider value={{ user, login, logout, loading }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)
