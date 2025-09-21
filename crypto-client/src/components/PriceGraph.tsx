'use client'
import { useQuery } from '@tanstack/react-query'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts'
import { fetchPriceHistory } from '@/lib/api'



export default function PriceGraph({ coinId }: { coinId: string }) {


  const { data, isLoading } = useQuery({
    queryKey: ['price-history', coinId],
    queryFn: () => fetchPriceHistory(coinId),
  })

  if (isLoading) return <p>Loading graph...</p>

  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={data}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="timestamp" />
        <YAxis />
        <Tooltip />
        <Line type="monotone" dataKey="price" stroke="#8884d8" />
      </LineChart>
    </ResponsiveContainer>
  )
}