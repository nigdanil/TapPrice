import { AppBar, Toolbar, Typography } from '@mui/material'

export default function Header() {
  return (
    <AppBar position="static" color="default" elevation={1}>
      <Toolbar>
        <Typography variant="h6" color="inherit" component="div">
          TapPrice Admin
        </Typography>
      </Toolbar>
    </AppBar>
  )
}
