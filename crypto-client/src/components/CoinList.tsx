import Link from 'next/link'
import coins from '../public/data/coins.js'
import { useQuery } from '@tanstack/react-query'

export default function CoinList() {
  

  return (
    <ul className="grid grid-cols-1 md:grid-cols-3 gap-4">
      {coins?.map((coin: { name: string }) => (
        <li key={coin.name} className="border p-4 rounded">
          <Link href={`/coin/${coin.name}`}>{coin.name}</Link>
        </li>
      ))}
    </ul>
  )
}