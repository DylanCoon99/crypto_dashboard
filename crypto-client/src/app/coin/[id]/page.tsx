'use client'
import { useState } from 'react'
import { useParams } from 'next/navigation';
import PriceGraph from '@/components/PriceGraph'
import AIInsights from '@/components/AIInsights'

export default function CoinPage() {
  const params = useParams();
  const coinId = params.id as string;
  const [showInsights, setShowInsights] = useState(false)

  


  return (
    <main className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4 capitalize">{coinId} Dashboard</h1>
      <p className="text-xl mb-4">Current Price: Loading...</p>  {/* Fetch current price */}
      <PriceGraph coinId={coinId} />
      <button 
        className="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
        onClick={() => setShowInsights(!showInsights)}
      >
        {showInsights ? 'Hide AI Insights' : 'Show AI Insights'}
      </button>
      {showInsights && <AIInsights coinId={coinId} />}
    </main>
  )
}