import { AppBar, Toolbar, Typography, Button } from '@mui/material'
import { useAuth } from '../auth/AuthContext'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

export default function Header() {
  const { logout } = useAuth()
  const navigate = useNavigate()

  const handleLogout = async () => {
    await axios.get('/api/logout', { withCredentials: true })
    logout()
    navigate('/login')
  }

  return (
    <AppBar position="static" color="default">
      <Toolbar>
        <Typography variant="h6" sx={{ flexGrow: 1 }}>
          TapPrice Admin
        </Typography>
        <Button onClick={handleLogout} color="inherit">
          Выйти
        </Button>
      </Toolbar>
    </AppBar>
  )
}
