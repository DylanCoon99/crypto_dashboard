'use client'  // Client-side for interactivity
import CoinList from '@/components/CoinList'

export default function Home() {
  return (
    <main className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4">Crypto Dashboard</h1>
      <CoinList />
    </main>
  )
}