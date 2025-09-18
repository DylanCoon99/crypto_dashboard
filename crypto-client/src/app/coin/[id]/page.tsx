'use client'
import { useState } from 'react'
import PriceGraph from '@/components/PriceGraph'
import AIInsights from '@/components/AIInsights'

export default function CoinPage({ params }: { params: { id: string } }) {
  const [showInsights, setShowInsights] = useState(false)

  return (
    <main className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4 capitalize">{params.id} Dashboard</h1>
      <p className="text-xl mb-4">Current Price: Loading...</p>  {/* Fetch current price */}
      <PriceGraph coinId={params.id} />
      <button 
        className="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
        onClick={() => setShowInsights(!showInsights)}
      >
        {showInsights ? 'Hide AI Insights' : 'Show AI Insights'}
      </button>
      {showInsights && <AIInsights coinId={params.id} />}
    </main>
  )
}