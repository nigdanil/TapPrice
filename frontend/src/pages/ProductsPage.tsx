// D:\TapPrice\frontend\src\pages\ProductsPage.tsx
import { useEffect, useState } from 'react'
import {
  Box, Button, Typography, Select, MenuItem, InputLabel, FormControl, Stack
} from '@mui/material'
import axios from 'axios'

export default function ProductsPage() {
  const [products, setProducts] = useState([])
  const [venues, setVenues] = useState([])
  const [categories, setCategories] = useState([])
  const [selectedVenue, setSelectedVenue] = useState('')
  const [selectedCategory, setSelectedCategory] = useState('')

  const fetchData = async () => {
    try {
      const res = await axios.get('/api/products', {
        params: {
          venue_id: selectedVenue || undefined,
          category_id: selectedCategory || undefined,
        },
        withCredentials: true,
      })
      setProducts(res.data)
    } catch (err) {
      console.error('Ошибка загрузки товаров', err)
    }
  }

  const fetchFilters = async () => {
    try {
      const [venuesRes, categoriesRes] = await Promise.all([
        axios.get('/api/venues'),
        axios.get('/api/categories'),
      ])
      setVenues(venuesRes.data)
      setCategories(categoriesRes.data)
    } catch (err) {
      console.error('Ошибка загрузки фильтров', err)
    }
  }

  useEffect(() => {
    fetchFilters()
  }, [])

  useEffect(() => {
    fetchData()
  }, [selectedVenue, selectedCategory])

  return (
    <Box>
      <Typography variant="h5" gutterBottom>Товары</Typography>

      <Stack direction="row" spacing={2} sx={{ mb: 2 }}>
        <FormControl sx={{ minWidth: 200 }}>
          <InputLabel>Площадка</InputLabel>
          <Select
            value={selectedVenue}
            label="Площадка"
            onChange={(e) => setSelectedVenue(e.target.value)}
          >
            <MenuItem value="">Все</MenuItem>
            {venues.map((v: any) => (
              <MenuItem key={v.id} value={v.id}>{v.name}</MenuItem>
            ))}
          </Select>
        </FormControl>

        <FormControl sx={{ minWidth: 200 }}>
          <InputLabel>Категория</InputLabel>
          <Select
            value={selectedCategory}
            label="Категория"
            onChange={(e) => setSelectedCategory(e.target.value)}
          >
            <MenuItem value="">Все</MenuItem>
            {categories.map((c: any) => (
              <MenuItem key={c.id} value={c.id}>{c.name}</MenuItem>
            ))}
          </Select>
        </FormControl>

        <Button variant="contained">Импорт</Button>
        <Button variant="outlined">Добавить товар</Button>
      </Stack>

      {/* Временно: простой список */}
      <pre>{JSON.stringify(products, null, 2)}</pre>
    </Box>
  )
}
