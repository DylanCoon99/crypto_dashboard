import Link from 'next/link'
import { useQuery } from '@tanstack/react-query'
import { fetchCoins } from '@/lib/api'

export default function CoinList() {
  const { data, isLoading } = useQuery({ queryKey: ['coins'], queryFn: fetchCoins })

  if (isLoading) return <p>Loading...</p>

  return (
    <ul className="grid grid-cols-1 md:grid-cols-3 gap-4">
      {data?.map((coin: { id: string; name: string }) => (
        <li key={coin.id} className="border p-4 rounded">
          <Link href={`/coin/${coin.id}`}>{coin.name}</Link>
        </li>
      ))}
    </ul>
  )
}