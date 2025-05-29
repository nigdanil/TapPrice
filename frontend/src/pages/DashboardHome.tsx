import { Typography, Box, Paper } from '@mui/material'

export default function DashboardHome() {
  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Добро пожаловать в TapPrice Admin
      </Typography>

      <Paper sx={{ p: 3, mt: 2 }}>
        <Typography variant="body1">
          Здесь будет сводка по данным: количество товаров, категорий, пользователей и активности.
        </Typography>
      </Paper>
    </Box>
  )
}
