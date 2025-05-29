import { Box, List, ListItemButton, ListItemText, Toolbar } from '@mui/material'
import { NavLink } from 'react-router-dom'

const links = [
  { label: 'Товары', path: '/dashboard/products' },
  { label: 'Категории', path: '/dashboard/categories' },
  { label: 'Площадки', path: '/dashboard/venues' },
  { label: 'Пользователи', path: '/dashboard/users' },
  { label: 'Аудит-лог', path: '/dashboard/audit-log' },
]

export default function Sidebar() {
  return (
    <Box sx={{ width: 240, bgcolor: 'grey.100', height: '100vh' }}>
      <Toolbar />
      <List>
        {links.map(({ label, path }) => (
          <ListItemButton
            component={NavLink}
            to={path}
            key={path}
            sx={{ '&.active': { bgcolor: 'primary.light', color: 'white' } }}
          >
            <ListItemText primary={label} />
          </ListItemButton>
        ))}
      </List>
    </Box>
  )
}
