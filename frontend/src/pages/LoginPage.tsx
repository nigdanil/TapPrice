// D:\TapPrice\frontend\src\pages\LoginPage.tsx
import { useForm } from 'react-hook-form'
import { Box, TextField, Button, Typography } from '@mui/material'
import { useAuth } from '../auth/AuthContext'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

export default function LoginPage() {
  const { register, handleSubmit } = useForm()
  const { login } = useAuth()
  const navigate = useNavigate()

  const onSubmit = async (data: any) => {
    try {
      // Отправляем логин/пароль, сервер ставит JWT в cookie
      await axios.post('/api/login', data, { withCredentials: true })

      // Сохраняем пользователя вручную (пока без роли)
      login({ username: data.username, role: 'unknown' })

      // Переход в админку
      navigate('/dashboard')
    } catch (err) {
      alert('Ошибка входа: неверный логин или пароль')
    }
  }

  return (
    <Box sx={{ maxWidth: 400, mx: 'auto', mt: 10 }}>
      <Typography variant="h5" gutterBottom>Вход</Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <TextField fullWidth label="Username" margin="normal" {...register('username')} />
        <TextField fullWidth label="Password" type="password" margin="normal" {...register('password')} />
        <Button type="submit" fullWidth variant="contained">Войти</Button>
      </form>
    </Box>
  )
}
